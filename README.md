# rotary phone

Прототип анализа блокировки мессенджеров. 

## Реализовано

- discord
  - тест  web
  - тест скачивания видео
- facebook
  - тест web
- instagram
  - тест web
- signal
  - web тест
- telegram
  - web тест
  - тест скачивания видео
- viber
  - web тест
- vk
  - web тест
  - тест скачивания видео
- whatsapp
  - web тест
- youtube
  - web тест
  - тест скачивания видео
  - тест доступа к thumbnail канала

## usage
тест мессенджеров (без скачивание видео)
```shell
docker-compose up

```
тест с скачиванием видео. Требуется python, yt-dlp

```shell
pip install yt-dlp
```

```shell
# docker without access root
sudo ./run_compose_a_video.sh
```