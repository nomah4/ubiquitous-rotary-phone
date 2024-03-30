# системные требования
Потенциальные требования нод

- linux
- умеет в docker (или поставлять bash скрипт установки. Компоненты желательно должны дружить с arm)
- если времени хватит - android (arm)
# Какие мессенджеры надо исследовать
## Базовые
- Telegram
- Signal
- WhatsApp
- Viber
## Дополнительно (пригодится для контрольной проверки):
- Facebook
- VK
- instagram
- discord

# методология исследования доступа

# Способы проверки:
- PING на исследуемый источник
- публичные/закрытые API. Способы отправки запроса по приоритету:
	- HEAD запрос и проверка по заголовкам для экономии трафика
	- GET
	- POST
- запрос на публичные приглашения в чаты (если мессенджер поддерживает)
- запрос на публичные страницы на группы, медиа контент, профили
- запрос на страницы разработчиков или популярных личностей
Например:
- телеграм - канал павла дурова
- вк - страница ВК дурова или технические id страницы до 100
- инстаграмм - профиль криштиану рональдо (или любой другой, который гарантированно не удалят)
## Если протокол отличается от HTTP rest-api
Например, если это signal или telegram - взять готовую реализацию протокола на любом доступном ЯП (приоритет - поддержка arm архитектуры).
- отправлять PING запрос (если возможно)
- отправлять умышленно неправильные данные авторизации и проверять результат/типовую ошибку библиотеки

# Мессенджеры
## Telegram
### ping

```
# telegram.org domain not support PING
ping -c 4 api.telegram.org
```

### main
```
curl -I https://telegram.org/

HTTP/2 200 
server: nginx/1.18.0
date: Fri, 29 Mar 2024 20:07:24 GMT
content-type: text/html; charset=utf-8
content-length: 19594
set-cookie: stel_ssid=...; expires=Sat, 30 Mar 2024 07:14:04 GMT; path=/; samesite=None; secure; HttpOnly
pragma: no-cache
cache-control: no-store
x-frame-options: SAMEORIGIN
strict-transport-security: max-age=31536000; includeSubDomains; preload
```
### web

```
curl -I https://web.telegram.org

HTTP/2 200 
server: nginx/1.18.0
date: Fri, 29 Mar 2024 19:51:45 GMT
content-type: text/html
content-length: 1587
last-modified: Wed, 25 Oct 2023 19:17:17 GMT
etag: "653969bd-633"
expires: Fri, 29 Mar 2024 20:51:45 GMT
cache-control: max-age=3600
x-frame-options: deny

```
### bot API

```
curl -I https://api.telegram.org

HTTP/2 302 
server: nginx/1.18.0
date: Fri, 29 Mar 2024 19:48:08 GMT
content-type: text/html
content-length: 145
location: https://core.telegram.org/bots
strict-transport-security: max-age=31536000; includeSubDomains; preload
access-control-allow-origin: *
access-control-allow-methods: GET, POST, OPTIONS
access-control-expose-headers: Content-Length,Content-Type,Date,Server,Connection
```
### channel
```
curl -I https://t.me/rove

HTTP/2 200 
server: nginx/1.18.0
date: Fri, 29 Mar 2024 19:52:19 GMT
content-type: text/html; charset=utf-8
content-length: 11052
set-cookie: stel_ssid=b918a6b943c26812a4_16066502942867410517; expires=Sat, 30 Mar 2024 19:52:19 GMT; path=/; samesite=None; secure; HttpOnly
pragma: no-cache
cache-control: no-store
x-frame-options: ALLOW-FROM https://web.telegram.org
content-security-policy: frame-ancestors https://web.telegram.org
strict-transport-security: max-age=35768000
```

### channel post
```
curl -I https://t.me/rove/1

HTTP/2 200 
server: nginx/1.18.0
date: Fri, 29 Mar 2024 19:50:44 GMT
content-type: text/html; charset=utf-8
content-length: 11136
set-cookie: stel_ssid=...; expires=Sat, 30 Mar 2024 19:50:44 GMT; path=/; samesite=None; secure; HttpOnly
pragma: no-cache
cache-control: no-store
x-frame-options: ALLOW-FROM https://web.telegram.org
content-security-policy: frame-ancestors https://web.telegram.org
strict-transport-security: max-age=35768000


```
### invite

> TODO найти стабильнее ссылку (это первая попавшееся)

```
curl -I https://t.me/+HfQFnx6tv3FlODRh

HTTP/2 200 
server: nginx/1.18.0
date: Fri, 29 Mar 2024 19:55:50 GMT
content-type: text/html; charset=utf-8
content-length: 11082
set-cookie: stel_ssid=...; expires=Sat, 30 Mar 2024 19:55:50 GMT; path=/; samesite=None; secure; HttpOnly
pragma: no-cache
cache-control: no-store
x-frame-options: ALLOW-FROM https://web.telegram.org
content-security-policy: frame-ancestors https://web.telegram.org
strict-transport-security: max-age=35768000

```
### client
использовать готовые реализации mtproto протоколов и оберток (использовать неправильные данные для авторизации)

- python https://github.com/pyrogram/pyrogram
- golang https://github.com/gotd/td
- возможно есть еще реализации mtproto
## WhatsApp
### ping
```
ping -c 4 whatsapp.com
```
### main page
```
curl -I https://www.whatsapp.com/

HTTP/2 200 
vary: Accept-Encoding
set-cookie: wa_lang_pref=ru; expires=Fri, 05-Apr-2024 20:05:42 GMT; Max-Age=604800; path=/; domain=.whatsapp.com; secure
set-cookie: wa_ul=c1da9614-87d9-41ba-86e9-60a944eb90a8; expires=Thu, 27-Jun-2024 20:05:42 GMT; Max-Age=7776000; path=/; domain=.www.whatsapp.com; secure; httponly
set-cookie: wa_csrf=EcB_YABdxoj8Cu32H309-R; path=/; domain=.whatsapp.com; secure; httponly
reporting-endpoints: coop_report="https://www.facebook.com/browser_reporting/coop/?minimize=0", coep_report="https://www.facebook.com/browser_reporting/coep/?minimize=0", default="https://www.whatsapp.com/whatsapp_browser_error_reports/?device_level=unknown", permissions_policy="https://www.whatsapp.com/whatsapp_browser_error_reports/"
report-to: {"max_age":2592000,"endpoints":[{"url":"https:\/\/www.facebook.com\/browser_reporting\/coop\/?minimize=0"}],"group":"coop_report","include_subdomains":true}, {"max_age":86400,"endpoints":[{"url":"https:\/\/www.facebook.com\/browser_reporting\/coep\/?minimize=0"}],"group":"coep_report"}, {"max_age":259200,"endpoints":[{"url":"https:\/\/www.whatsapp.com\/whatsapp_browser_error_reports\/?device_level=unknown"}]}, {"max_age":21600,"endpoints":[{"url":"https:\/\/www.whatsapp.com\/whatsapp_browser_error_reports\/"}],"group":"permissions_policy"}
content-security-policy: default-src 'self' data: blob:;script-src 'self' data: blob: *.whatsapp.com *.whatsapp.net *.twitter.com *.facebook.com *.facebook.net 'unsafe-inline' 'unsafe-eval';style-src 'self' data: blob: *.whatsapp.com *.whatsapp.net 'unsafe-inline' *.facebook.com;connect-src 'self' data: blob: *.whatsapp.com *.whatsapp.net wss://*.facebook.com:* *.fbcdn.net;font-src data: *.whatsapp.com *.whatsapp.net *.facebook.com static.xx.fbcdn.net https://fonts.gstatic.com;img-src 'self' data: blob: *.whatsapp.com *.whatsapp.net *.facebook.com *.fbcdn.net static.xx.fbcdn.net *.ytimg.com *.twitter.com;media-src 'self' data: blob: *.fbcdn.net;frame-src 'self' data: blob: *.twitter.com *.facebook.com *.youtube-nocookie.com *.youtube.com *.whatsapp.com;block-all-mixed-content;upgrade-insecure-requests;
permissions-policy: accelerometer=(), ambient-light-sensor=(), attribution-reporting=(), autoplay=(), bluetooth=(), camera=(), ch-device-memory=(), ch-save-data=(), ch-ua-arch=(), ch-ua-bitness=(), clipboard-read=(), clipboard-write=(), display-capture=(), encrypted-media=(), fullscreen=(self), gamepad=(), geolocation=(), gyroscope=(), hid=(), idle-detection=(), keyboard-map=(), local-fonts=(), magnetometer=(), microphone=(), midi=(), otp-credentials=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), screen-wake-lock=(), serial=(), usb=(), window-management=(), xr-spatial-tracking=();report-to="permissions_policy"
cross-origin-resource-policy: same-origin
cross-origin-embedder-policy-report-only: require-corp;report-to="coep_report"
cross-origin-opener-policy: same-origin-allow-popups;report-to="coop_report"
pragma: no-cache
cache-control: private, no-cache, no-store, must-revalidate
expires: Sat, 01 Jan 2000 00:00:00 GMT
x-content-type-options: nosniff
x-xss-protection: 0
x-frame-options: DENY
strict-transport-security: max-age=31536000; preload; includeSubDomains
content-type: text/html; charset="utf-8"
x-fb-debug: ...
date: Fri, 29 Mar 2024 20:05:42 GMT
x-fb-connection-quality: EXCELLENT; q=0.9, rtt=29, rtx=0, c=16, mss=1380, tbw=3468, tp=-1, tpl=-1, uplat=388, ullat=0
alt-svc: h3=":443"; ma=86400

```
### web
```
curl -I https://web.whatsapp.com/
HTTP/2 200 
content-type: text/html;charset=utf-8
content-length: 3640
vary: Accept-Encoding, User-Agent, Accept-Language
vary: Accept-Encoding
reporting-endpoints: coop_report="https://www.facebook.com/browser_reporting/coop/?minimize=0", coep_report="https://www.facebook.com/browser_reporting/coep/?minimize=0", default="https://www.whatsapp.com/whatsapp_browser_error_reports/?device_level=unknown", permissions_policy="https://www.whatsapp.com/whatsapp_browser_error_reports/"
report-to: {"max_age":2592000,"endpoints":[{"url":"https:\/\/www.facebook.com\/browser_reporting\/coop\/?minimize=0"}],"group":"coop_report","include_subdomains":true}, {"max_age":86400,"endpoints":[{"url":"https:\/\/www.facebook.com\/browser_reporting\/coep\/?minimize=0"}],"group":"coep_report"}, {"max_age":259200,"endpoints":[{"url":"https:\/\/www.whatsapp.com\/whatsapp_browser_error_reports\/?device_level=unknown"}]}, {"max_age":21600,"endpoints":[{"url":"https:\/\/www.whatsapp.com\/whatsapp_browser_error_reports\/"}],"group":"permissions_policy"}
document-policy: force-load-at-top
permissions-policy: accelerometer=(), ambient-light-sensor=(), attribution-reporting=(), autoplay=*, bluetooth=(), camera=(self), ch-device-memory=(), ch-save-data=(), ch-ua-arch=(), ch-ua-bitness=(), clipboard-read=(), clipboard-write=(self), display-capture=(), encrypted-media=(), fullscreen=(self), gamepad=(), geolocation=(), gyroscope=(), hid=(), idle-detection=(), keyboard-map=(), local-fonts=(), magnetometer=(), microphone=(self), midi=(), otp-credentials=(), payment=(), picture-in-picture=*, publickey-credentials-get=(), screen-wake-lock=(), serial=(), usb=(), window-management=(), xr-spatial-tracking=();report-to="permissions_policy"
cross-origin-resource-policy: cross-origin
cross-origin-embedder-policy-report-only: require-corp;report-to="coep_report"
cross-origin-opener-policy: unsafe-none;report-to="coop_report"
pragma: no-cache
cache-control: no-cache
expires: Sat, 01 Jan 2000 00:00:00 GMT
x-content-type-options: nosniff
x-xss-protection: 0
content-security-policy: frame-ancestors 'self';
content-security-policy: default-src 'self' data: blob:;script-src 'self' data: blob: 'unsafe-eval' 'unsafe-inline' 'report-sample' https://static.whatsapp.net https://*.youtube.com https://maps.googleapis.com https://maps.gstatic.com https://*.google-analytics.com;style-src 'self' data: blob: 'unsafe-inline' https://static.whatsapp.net https://fonts.googleapis.com;connect-src 'self' data: blob: https://*.whatsapp.net https://www.facebook.com https://crashlogs.whatsapp.net/wa_clb_data https://crashlogs.whatsapp.net/wa_fls_upload_check wss://*.web.whatsapp.com wss://web.whatsapp.com wss://web-fallback.whatsapp.com https://www.whatsapp.com https://dyn.web.whatsapp.com https://graph.whatsapp.com/graphql/ https://graph.facebook.com/graphql https://*.tenor.co https://*.giphy.com https://maps.googleapis.com https://*.google-analytics.com;font-src data: 'self' https://static.whatsapp.net fonts.googleapis.com https://fonts.gstatic.com;img-src 'self' data: blob: https://*.whatsapp.net https://*.fbcdn.net *.tenor.co *.tenor.com *.giphy.com https://*.ytimg.com https://maps.googleapis.com/maps/api/staticmap https://*.google-analytics.com;media-src 'self' data: blob: https://*.whatsapp.net https://*.cdninstagram.com https://*.fbcdn.net mediastream: *.tenor.co *.tenor.com https://*.giphy.com;child-src 'self' data: blob:;frame-src 'self' data: blob: https://www.youtube.com/embed/;block-all-mixed-content;upgrade-insecure-requests;
strict-transport-security: max-age=63072000; includeSubDomains; preload
x-fb-debug: ...
date: Fri, 29 Mar 2024 19:57:32 GMT
x-fb-connection-quality: EXCELLENT; q=0.9, rtt=30, rtx=0, c=16, mss=1380, tbw=3469, tp=-1, tpl=-1, uplat=43, ullat=0
alt-svc: h3=":443"; ma=86400

```
### invite

> TODO create stable link. this found by youtube search query `whatsapp pubg` + `live` filter


```
# https://www.youtube.com/watch?v=U-SMUOHAuyg
curl -I https://chat.whatsapp.com/K5wu27LINloHgjucK0rSZy

HTTP/2 200 
vary: Accept-Encoding
set-cookie: wa_lang_pref=ru; expires=Fri, 05-Apr-2024 20:02:22 GMT; Max-Age=604800; path=/; domain=.whatsapp.com; secure
reporting-endpoints: coop_report="https://www.facebook.com/browser_reporting/coop/?minimize=0", coep_report="https://www.facebook.com/browser_reporting/coep/?minimize=0", default="https://www.whatsapp.com/whatsapp_browser_error_reports/?device_level=unknown", permissions_policy="https://www.whatsapp.com/whatsapp_browser_error_reports/"
report-to: {"max_age":2592000,"endpoints":[{"url":"https:\/\/www.facebook.com\/browser_reporting\/coop\/?minimize=0"}],"group":"coop_report","include_subdomains":true}, {"max_age":86400,"endpoints":[{"url":"https:\/\/www.facebook.com\/browser_reporting\/coep\/?minimize=0"}],"group":"coep_report"}, {"max_age":259200,"endpoints":[{"url":"https:\/\/www.whatsapp.com\/whatsapp_browser_error_reports\/?device_level=unknown"}]}, {"max_age":21600,"endpoints":[{"url":"https:\/\/www.whatsapp.com\/whatsapp_browser_error_reports\/"}],"group":"permissions_policy"}
content-security-policy: default-src 'self' data: blob:;script-src 'self' data: blob: 'unsafe-eval' 'unsafe-inline' *.facebook.com *.fbcdn.net *.whatsapp.com *.whatsapp.net https://*.facebook.net;style-src 'self' data: blob: 'unsafe-inline' https://fonts.googleapis.com *;connect-src 'self' data: blob: https://*.whatsapp.com;font-src data: https://*.fbcdn.net https://static.whatsapp.net;img-src 'self' data: blob: *;frame-src 'self' data: blob: whatsapp:;block-all-mixed-content;upgrade-insecure-requests;
permissions-policy: accelerometer=(), ambient-light-sensor=(), attribution-reporting=(), autoplay=(), bluetooth=(), camera=(), ch-device-memory=(), ch-save-data=(), ch-ua-arch=(), ch-ua-bitness=(), clipboard-read=(), clipboard-write=(), display-capture=(), encrypted-media=(), fullscreen=(self), gamepad=(), geolocation=(), gyroscope=(), hid=(), idle-detection=(), keyboard-map=(), local-fonts=(), magnetometer=(), microphone=(), midi=(), otp-credentials=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), screen-wake-lock=(), serial=(), usb=(), window-management=(), xr-spatial-tracking=();report-to="permissions_policy"
cross-origin-resource-policy: cross-origin
cross-origin-embedder-policy-report-only: require-corp;report-to="coep_report"
cross-origin-opener-policy: same-origin;report-to="coop_report"
pragma: no-cache
cache-control: private, no-cache, no-store, must-revalidate
expires: Sat, 01 Jan 2000 00:00:00 GMT
x-content-type-options: nosniff
x-xss-protection: 0
x-frame-options: DENY
strict-transport-security: max-age=31536000; preload; includeSubDomains
content-type: text/html; charset="utf-8"
x-fb-debug: ...
date: Fri, 29 Mar 2024 20:02:22 GMT
x-fb-connection-quality: EXCELLENT; q=0.9, rtt=30, rtx=0, c=16, mss=1380, tbw=3469, tp=-1, tpl=-1, uplat=590, ullat=0
alt-svc: h3=":443"; ma=86400


```
## Viber
### ping

```
ping -c 4 viber.com
```

### invite

> google: viber invite 

> Инвайт на Объясняем.рф группу

```
curl -I "https://invite.viber.com/?g2=AQAp6MUgQTByH07Ng7dOBCcs2Gric2EkPLB9qRIGDhnhC%2FUT%2Fz7kBw2hKPwTArUs&lang=en"
HTTP/2 200 
content-type: text/html; charset=utf-8
server: nginx
x-powered-by: Express
etag: W/"af09-XNwky1jGV+M/O8NtAZQPYiJHQ48"
cache-control: max-age=10141
expires: Fri, 29 Mar 2024 23:30:29 GMT
date: Fri, 29 Mar 2024 20:41:28 GMT
```
### main
```
curl -I https://viber.com
HTTP/2 301 
server: AkamaiGHost
content-length: 0
location: https://www.viber.com/
cache-control: max-age=0
expires: Fri, 29 Mar 2024 20:39:08 GMT
date: Fri, 29 Mar 2024 20:39:08 GMT
server-timing: ak_p; desc="1711744748691_390846690_938140360_9_336_93_55_15";dur=1
```

### channel post

```
curl -I "https://invite.viber.com/?g2=AQAp6MUgQTByH07Ng7dOBCcs2Gric2EkPLB9qRIGDhnhC%2FUT%2Fz7kBw2hKPwTArUs&lang=en&mi=11033"
HTTP/2 200 
content-type: text/html; charset=utf-8
server: nginx
x-powered-by: Express
etag: W/"aef8-Mi2YHHtUs3G//kHROlSJVGhyL4I"
cache-control: max-age=10411
expires: Fri, 29 Mar 2024 23:36:20 GMT
date: Fri, 29 Mar 2024 20:42:49 GMT

```
## VK

> HEAD requests returns 418 (im a teapot)
### page

```
curl -I https://vk.com/id1
HTTP/2 418 
server: kittenx
date: Fri, 29 Mar 2024 21:11:18 GMT
content-length: 0
x-frontend: front656402
access-control-expose-headers: X-Frontend
x-trace-id: lABgiBOTMtH_l_biDCv-b2c_IVujiQ

```
### ping

```
ping -c 4 vk.com
```
### main

```
curl -I vk.com
HTTP/1.1 418 
Server: kittenx
Date: Fri, 29 Mar 2024 20:10:00 GMT
Content-Length: 0
Connection: keep-alive
X-Frontend: front656600
Access-Control-Expose-Headers: X-Frontend
X-Trace-Id: xUKe9WArmzWStOWtRqJw4UuMlhRrrw

```

### bot api

```
curl -I https://api.vk.com/method/
HTTP/2 418 
server: kittenx
date: Fri, 29 Mar 2024 20:11:36 GMT
content-length: 0
x-trace-id: 14aAa9bFb91LLVjzUgNTprenVPmKQQ
```

```
# GET
curl https://api.vk.com/method/
{"error": {"error_code": 3,"error_msg": "Not found"}}
```

### auth

> НАСТОЯЩИЕ ДАННЫЕ ИСПОЛЬЗОВАТЬ НЕ РЕКОМЕНДУЕТСЯ ДЛЯ ПРОДАКШЕНА, ВК МОЖЕТ ЗАМОРОЗИТЬ СТРАНИЦУ ЗА IMPLICT FLOW АВТОРИЗАЦИЮ

```python
import requests
s = requests.Session()

s.get('https://oauth.vk.com/token', 
	  params={"grant_type":"password", "client_id":2274003, "username":email,"password":password, "client_secret":"hHbZxrka2uZ6jB1inYsH"}
```

```
curl -I "https://oauth.vk.com/token?grant_type=password&client_id=2274003&username=88005553535&password=12345678&client_secret=hHbZxrka2uZ6jB1inYsH"
HTTP/2 418 
server: kittenx
date: Fri, 29 Mar 2024 21:01:13 GMT
content-length: 0
strict-transport-security: max-age=15768000
x-trace-id: wG0aJmMkXm9UjhpkRPOa8jIK8IEHSg
```

```
curl "https://oauth.vk.com/token?grant_type=password&client_id=2274003&username=88005553535&password=12345678&client_secret=hHbZxrka2uZ6jB1inYsH"
{"error":"need_captcha","captcha_sid":"153072126722","is_refresh_enabled":true,"captcha_img":"https:\/\/vk.com\/captcha.php?sid=153072126722&source=api-oauth&app_id=2274003&device_id=&resized=1","captcha_ts":1711746107.687000,"captcha_attempt":1,"captcha_ratio":2.600000,"is_sound_captcha_available":true,"captcha_track":"https:\/\/vk.com\/sound_captcha.php?captcha_sid=153072126722&act=get&source=api-oauth&app_id=2274003&device_id=","is_enabled_vkui":true,"uiux_changes":false}
```
## instagram

> TODO
## Facebook

> TODO

## discord

### ping

```
ping -c 4 discord.com
```
### main

```
curl -I https://discord.com
HTTP/2 200 
date: Fri, 29 Mar 2024 20:33:45 GMT
content-type: text/html
cf-ray: 86c2a305bb759d7a-DME
cf-cache-status: HIT
cache-control: no-cache
last-modified: Fri, 29 Mar 2024 20:33:36 GMT
set-cookie: __dcfduid=...; Expires=Wed, 28 Mar 2029 20:33:45 GMT; Max-Age=157680000; Path=/; Secure; HttpOnly; SameSite=Lax
strict-transport-security: max-age=31536000; includeSubDomains; preload
vary: Accept-Encoding,x-wf-forwarded-proto
content-security-policy: default-src 'self'; script-src 'self' 'unsafe-inline' 'nonce-OSwxMDAsNjEsMjQ0LDEwNywzOSwzOCwyMTY=' https://discord.com https://www.googletagmanager.com https://connect.facebook.net https://www.google-analytics.com https://ssl.google-analytics.com https://www.gstatic.com/recaptcha/ https://www.google.com/recaptcha/ https://recaptcha.net/recaptcha/ https://hcaptcha.com https://*.hcaptcha.com https://s.ytimg.com/yts/jsbin/ https://www.youtube.com/iframe_api https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location https://script.crazyegg.com https://*.website-files.com https://global.localizecdn.com https://d3e54v103j8qbb.cloudfront.net https://gist.github.com https://unpkg.com/@splinetool/runtime/build/runtime.js https://*.twitter.com https://cdn.jsdelivr.net/npm/slick-carousel@1.8.1/slick/slick.min.js https://boards-api.greenhouse.io https://ajax.googleapis.com/ajax/libs/webfont/1.6.26/webfont.js https://ajax.googleapis.com/ajax/libs/jquery/3.6.0/jquery.min.js https://cdn.finsweet.com/files/fscalendar/calendar-invite-v1.0.min.js 'sha256-mjdgHR9aXy-6OwAGlNS_XgNcYG1Uhd2U4pl8vi7-XCY=' 'sha256-gqG2LEZaHDwOL3S_CXJTuk_f3LimTEyruhOc_U0_QUY=' 'sha256-y0oGiuXZdmX7xRABTnY5cbHkfghDqbfX6JoerXLgVJc=' 'sha256-gBzDBwsujjXjXk6GLgdSlLIrvt5h0s3F_qF7Qt8TYwY=' 'sha256-TrY3AqlyKfZdsI3LYsy6u8GAhckLEXeyLcFK2gOe18U=' 'sha256-lVOL-gH47X0Li5QriWNZ69Hcr-71DsXFvGmQxN9TpBw=' 'sha256-j11ZNhk91nmUjPCBAIRcvJeEgnkbdJ9qNqoEMekilec=' 'sha256-1sQ9sTbc6Lumd2Frwf7IBwGG02gPTreTI8QBBW5kibM=' 'sha256-uh1p-Vy3_Cn66Ugk4Hak-gGr2Udg7yiI_5u5E_BdCRM=' 'sha256-7JHgDILwD7i_kvnHwJFF5WsHHmIc98tkBqDqbv47iFE=' 'sha256-KvstP_RIj6GGaE25Mqo-kIO0_WVEls1n5tnNhm8zmPA=' 'sha256-6xIDOlx5P0LKHv8fkot5ULOnB8ySdhjJi5r_ZP5EDPY=' 'sha256-jY_7jWrddtNUb-Y4CFKWaH-R2lrqgm_LAX72E8SLqKw=' 'sha256-MdICB9cW7ILT3ZeSxhN2YlpFxEsn5WHr03Ix-WVpHsw=' 'sha256-fUfByJGhChEFu7PE5HJfFwiYKySnP1H0iXvAxkauLNU=' 'sha256-xjkCDxBOM2TlIn5DpGQM4aJldb4AiHMKlRjfW46l-x0=' 'sha256-VOPfGBY-XgTDMwhG41S5eZyMKlu3gN60suwCPDWZ8MY=' 'sha256-tVeTMYknRG_IAdCHRGlDd9S2bX2_rX0e4HpaP9lgKWY=' 'sha256-kprfDg8ElCpUCFQAX5shnAPf3i59vVTSy02AjZXV3k0=' 'sha256-llLws8TR-U3nNRCIvJNVc-SGscqwyeO1IPgpbnWuZdc=' 'sha256-h9lm4cvrD7egZu1GTAE1h2IDy1K4fXgD-q_O7aEosuw=' 'sha256-_cdQbTQzcfSt2_aCceUvkUmLh1WMdvlKbi1BBG7u8Jg=' 'sha256-U0jHWhsvIpjnwYKeJS_-2pe9ROsYnck5ZB2aXNyKWq8=' 'sha256-rB4G_-e_bAPU7rKI_9HC1lBZ0XEa_nHDH6hXFz4GIh4=' 'sha256-N02bP-slnHB-OYEN6imRqCHcHLN5DvBouRmyO2qcQYU=' 'sha256-QHiY6i8ql9SJTaFXzUhm08ZWuNz0QarKruf0Omd9-OQ=' 'sha256-s4OBHcHJnkGxjEyNJhU5BQt4qlt6MH07rG/j/hFOUnE=' 'sha256-s4OBHcHJnkGxjEyNJhU5BQt4qlt6MH07rG_j_hFOUnE=' 'sha256-mjdgHR9aXy+6OwAGlNS/XgNcYG1Uhd2U4pl8vi7+XCY=' 'sha256-jY/7jWrddtNUb+Y4CFKWaH+R2lrqgm/LAX72E8SLqKw=' 'sha256-lVOL+gH47X0Li5QriWNZ69Hcr+71DsXFvGmQxN9TpBw=' 'sha256-/cdQbTQzcfSt2/aCceUvkUmLh1WMdvlKbi1BBG7u8Jg=' 'sha256-N02bP+slnHB+OYEN6imRqCHcHLN5DvBouRmyO2qcQYU=' 'sha256-gqG2LEZaHDwOL3S/CXJTuk/f3LimTEyruhOc/U0/QUY=' 'sha256-llLws8TR+U3nNRCIvJNVc+SGscqwyeO1IPgpbnWuZdc=' 'sha256-gBzDBwsujjXjXk6GLgdSlLIrvt5h0s3F/qF7Qt8TYwY=' 'sha256-6xIDOlx5P0LKHv8fkot5ULOnB8ySdhjJi5r/ZP5EDPY=' 'sha256-7JHgDILwD7i/kvnHwJFF5WsHHmIc98tkBqDqbv47iFE=' 'sha256-VOPfGBY+XgTDMwhG41S5eZyMKlu3gN60suwCPDWZ8MY='; style-src 'self' 'unsafe-inline' https://fonts.googleapis.com https://*.hcaptcha.com https://hcaptcha.com https://*.website-files.com https://*.githubassets.com; img-src 'self' https://www.google-analytics.com https://www.googletagmanager.com https://www.facebook.com https://cdn.discordapp.com https://hackerone-api.discord.workers.dev/user-avatars/ https://safety.discord.com https://discordmoderatoracademy.zendesk.com https://assets-global.website-files.com data: https://*.website-files.com https://global.localizecdn.com https://*.ytimg.com https://uploads-ssl.webflow.com; font-src 'self' https://fonts.gstatic.com https://fonts.gstatic.com https://*.website-files.com; connect-src 'self' https://discordapp.com https://discord.com https://connect.facebook.net https://api.greenhouse.io https://api.github.com https://sentry.io https://www.google-analytics.com https://hackerone-api.discord.workers.dev https://*.hcaptcha.com https://hcaptcha.com https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location ws://127.0.0.1:* http://127.0.0.1:* https://global.localizecdn.com https://*.website-files.com https://webflow.com/api/; media-src 'self' https://cdn.discordapp.com/assets/; frame-src https://discordapp.com/domain-migration https://www.google.com/recaptcha/ https://recaptcha.net/recaptcha/ https://*.hcaptcha.com https://hcaptcha.com https://www.youtube.com/embed/ https://hackerone.com/631fba12-9388-43c3-8b48-348f11a883c0/ https://10851314.fls.doubleclick.net/ https://*.twitter.com https://*.vimeo.com;
permissions-policy: interest-cohort=()
x-content-type-options: nosniff
x-frame-options: DENY
x-xss-protection: 1; mode=block
report-to: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v4?s=ykEOAGT1MvE%2FW%2BCNH6D6TqL9TvbpDKXW7ysKsjWfJMwxbKkzbpKX2w2ssjQDdDmQjMFUlkulvgMH%2FIx9ZEmNMqQ0vnf1YcJmyad%2BuWQ528CKrDtmYMuuaS%2FMEFZE"}],"group":"cf-nel","max_age":604800}
nel: {"success_fraction":0,"report_to":"cf-nel","max_age":604800}
set-cookie: __sdcfduid=...; Expires=Wed, 28 Mar 2029 20:33:45 GMT; Max-Age=157680000; Path=/; Secure; HttpOnly; SameSite=Lax
set-cookie: __cfruid=...; path=/; domain=.discord.com; HttpOnly; Secure; SameSite=None
set-cookie: _cfuvid=...; path=/; domain=.discord.com; HttpOnly; Secure; SameSite=None
server: cloudflare
alt-svc: h3=":443"; ma=86400

```
### bot api

```
curl -I https://discord.com/api/GetUser

HTTP/2 404 
date: Fri, 29 Mar 2024 20:32:35 GMT
content-type: application/json
content-length: 40
set-cookie: __dcfduid=...; Expires=Wed, 28-Mar-2029 20:32:35 GMT; Max-Age=157680000; Secure; HttpOnly; Path=/; SameSite=Lax
strict-transport-security: max-age=31536000; includeSubDomains; preload
via: 1.1 google
alt-svc: h3=":443"; ma=86400
cf-cache-status: DYNAMIC
report-to: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v4?s=byvLppH%2B6aTtMa5V5Hrhdsnh47AthH%2Fz4%2F71U2NliopPZvkqE77n0jeFMkXSyE1zfRVwXmOvPivdSKtZjQZgCPK5Yr5m%2FigN306vjCQo5%2Fw7O5%2FtqtF9oXuBdHeo"}],"group":"cf-nel","max_age":604800}
nel: {"success_fraction":0,"report_to":"cf-nel","max_age":604800}
x-content-type-options: nosniff
content-security-policy: frame-ancestors 'none'; default-src 'none'
set-cookie: __sdcfduid...; Expires=Wed, 28-Mar-2029 20:32:35 GMT; Max-Age=157680000; Secure; HttpOnly; Path=/; SameSite=Lax
set-cookie: __cfruid=...; path=/; domain=.discord.com; HttpOnly; Secure; SameSite=None
set-cookie: _cfuvid=...; path=/; domain=.discord.com; HttpOnly; Secure; SameSite=None
server: cloudflare
cf-ray: 86c2a14aed049e1b-DME

```

### invite

> NOTE: Random invite founded in youtube by `discord.gg` query + live filter

```
curl -I https://discord.com/invite/donut
HTTP/2 200 
date: Fri, 29 Mar 2024 20:35:41 GMT
content-type: text/html
cf-ray: 86c2a5d46c6d9df6-DME
cf-cache-status: HIT
accept-ranges: bytes
cache-control: private
last-modified: Fri, 29 Mar 2024 19:56:42 GMT
set-cookie: __dcfduid=...; Expires=Wed, 28 Mar 2029 20:35:41 GMT; Max-Age=157680000; Path=/; Secure; HttpOnly; SameSite=Lax
strict-transport-security: max-age=31536000; includeSubDomains; preload
content-security-policy: default-src 'self'; script-src 'self' 'unsafe-eval' 'unsafe-inline' 'nonce-MjE4LDExOSwyMTUsNDgsMTMxLDEwMyw2OCwyNTE=' blob: https://cdn.discordapp.com/animations/ https://www.gstatic.com/recaptcha/ https://www.google.com/recaptcha/ https://recaptcha.net/recaptcha/ https://*.hcaptcha.com https://hcaptcha.com https://js.stripe.com https://js.braintreegateway.com https://assets.braintreegateway.com https://www.paypalobjects.com https://checkout.paypal.com https://c.paypal.com https://kit.cash.app; style-src 'self' 'unsafe-inline' https://cdn.discordapp.com https://*.hcaptcha.com https://hcaptcha.com https://kit.cash.app; img-src 'self' blob: data: https://*.discordapp.net https://*.discordapp.com https://*.discord.com https://i.scdn.co https://i.ytimg.com https://i.imgur.com https://media.tenor.co https://media.tenor.com https://c.tenor.com https://*.youtube.com https://*.giphy.com https://static-cdn.jtvnw.net https://pbs.twimg.com https://assets.braintreegateway.com https://checkout.paypal.com https://c.paypal.com https://b.stats.paypal.com https://slc.stats.paypal.com https://hnd.stats.paypal.com https://api.cash.app; font-src 'self' https://fonts.gstatic.com https://cash-f.squarecdn.com; connect-src 'self' https://status.discordapp.com https://status.discord.com https://support.discordapp.com https://support.discord.com https://discordapp.com https://discord.com https://discord-attachments-uploads-prd.storage.googleapis.com https://cdn.discordapp.com https://media.discordapp.net https://images-ext-1.discordapp.net https://images-ext-2.discordapp.net https://router.discordapp.net wss://*.discord.gg https://best.discord.media https://latency.discord.media wss://*.discord.media wss://dealer.spotify.com https://api.spotify.com https://music.amazon.com/embed/oembed https://sentry.io https://api.twitch.tv https://api.stripe.com https://api.braintreegateway.com https://client-analytics.braintreegateway.com https://*.braintree-api.com https://www.googleapis.com https://*.algolianet.com https://*.hcaptcha.com https://hcaptcha.com https://*.algolia.net ws://127.0.0.1:* http://127.0.0.1:*; media-src 'self' blob: disclip: https://*.discordapp.net https://*.discord.com https://*.discordapp.com https://*.youtube.com https://streamable.com https://vid.me https://twitter.com https://oddshot.akamaized.net https://*.giphy.com https://i.imgur.com https://media.tenor.co https://media.tenor.com https://c.tenor.com; frame-src https://discordapp.com/domain-migration discord: https://www.google.com/recaptcha/ https://recaptcha.net/recaptcha/ https://*.hcaptcha.com https://hcaptcha.com https://js.stripe.com https://hooks.stripe.com https://checkout.paypal.com https://c.paypal.com https://assets.braintreegateway.com https://checkoutshopper-live.adyen.com https://kit.cash.app https://player.twitch.tv https://clips.twitch.tv/embed https://player.vimeo.com https://www.youtube.com/embed/ https://www.tiktok.com/embed/ https://music.amazon.com/embed/ https://music.amazon.co.uk/embed/ https://music.amazon.de/embed/ https://music.amazon.co.jp/embed/ https://music.amazon.es/embed/ https://music.amazon.fr/embed/ https://music.amazon.it/embed/ https://music.amazon.com.au/embed/ https://music.amazon.in/embed/ https://music.amazon.ca/embed/ https://music.amazon.com.mx/embed/ https://music.amazon.com.br/embed/ https://www.youtube.com/s/player/ https://twitter.com/i/videos/ https://www.funimation.com/player/ https://www.redditmedia.com/mediaembed/ https://open.spotify.com/embed/ https://w.soundcloud.com/player/ https://audius.co/embed/ https://*.watchanimeattheoffice.com https://sessionshare.sp-int.playstation.com/embed/ https://localhost:* https://*.discordsays.com https://discordappcom.cloudflareaccess.com/; child-src 'self' blob: https://assets.braintreegateway.com https://checkout.paypal.com https://c.paypal.com; prefetch-src 'self' https://cdn.discordapp.com/assets/;
cross-origin-opener-policy: same-origin-allow-popups
permissions-policy: interest-cohort=()
x-build-id: f1e6efeefaedbc11f9e0fde8351e38a7bc0a126e
x-content-type-options: nosniff
x-frame-options: DENY
x-xss-protection: 1; mode=block
report-to: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v4?s=vDHKkxRU7HPSW2cQ1AfSDtToFSkqkxBk4lXws2R8Em0hP0G7RMh6lchM0BMNXhDExdwbsGNNBNwW8sKfDzu%2Fr4xgMN4E7fpmpIYJkh2raUgUSPazTYmlDYCD6GM5"}],"group":"cf-nel","max_age":604800}
nel: {"success_fraction":0,"report_to":"cf-nel","max_age":604800}
set-cookie: __sdcfduid=e8c6d5e1ee0b11ee8adbc92e53c4f0624f368d2d4eea7165a33899aaaeee69e49c1bd4a9210385c3484e0c240073a0e7; Expires=Wed, 28 Mar 2029 20:35:41 GMT; Max-Age=157680000; Path=/; Secure; HttpOnly; SameSite=Lax
set-cookie: __cfruid=...; path=/; domain=.discord.com; HttpOnly; Secure; SameSite=None
set-cookie: _cfuvid=...; path=/; domain=.discord.com; HttpOnly; Secure; SameSite=None
server: cloudflare
alt-svc: h3=":443"; ma=86400

```
### auth

Для авторизации пользователей необходимо пройти каптчу. Или заблаговременно извлекать токен (он не бесконечно живет)

## Signal

> мало обнаруженных ручек, требуется дополнительный ресерч

### ping

```
ping -c 4 signal.org
```

### main

```
curl -I https://signal.org
HTTP/2 200 
content-type: text/html; charset=utf-8
content-length: 5669
x-amz-id-2: UTLW5gEwxZdTldmQDbWAwicrzNlzQt0of0iWG9W8B3mb//vnqZ8GNt5hhbMk9FAfZkAWO8nc9OA=
x-amz-request-id: HNCKGJMFS8XDWZAE
date: Fri, 29 Mar 2024 20:49:57 GMT
content-encoding: gzip
last-modified: Fri, 15 Mar 2024 18:44:11 GMT
etag: "97a5f45438a9f2e32e5c46ca7aa4092b"
server: AmazonS3
strict-transport-security: max-age=31536000; preload
x-frame-options: DENY
x-cache: Miss from cloudfront
via: 1.1 11b1425a6d4f554d768315c2301c82b2.cloudfront.net (CloudFront)
x-amz-cf-pop: HEL51-P1
x-amz-cf-id: -KT5nY7LRDfx-8HZIDX5MkDXbhsEQmqGtndivcgqL9qPOzjhQA4I2Q==
```

### client
Реализовать неудачную авторизацию
- go https://github.com/bbernhard/signal-cli-rest-api
- rust https://github.com/signalapp/libsignal
- прочий поиск на ГХ