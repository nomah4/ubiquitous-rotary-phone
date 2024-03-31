# rotary phone

Прототип модулей для анализа блокировок мессенджеров.

> один мессенджер == 1 контейнер

## Demo

https://github.com/nomah4/ubiquitous-rotary-phone/assets/59173419/62483727-b132-4be2-8628-827366057059

## schema

```mermaid
flowchart TD
    SERVER("dpidetecor server")
    subgraph Messengers
        VK["VKCOM"]
        TG["VIBER"]
        INST["INSTAGRAM"]
        ABC["any"]
    end

    A[rotary phone] -->|launch containers| C{Run tests}

    C --> VK --> VK1
    C --> TG --> TG1
    C --> INST --> I1
    C .->|"Test cases"| ABC["Other messenger(s)"]
    
    subgraph TEST CASES 
        VK1["Test Response"] --> VK2["Test favicon sha256"]
        VK2 --> VK3["Test web"]
        VK3 --> VK4["Test video dl"]
        VK4 --> VK5["PING"]
        VK5 -->|"report"| SERVER
        
        TG1["Test Response"] --> TG2["Test favicon sha256"]
        TG2 --> TG3["Test web"]
        TG3 -->|"report"| SERVER
        
        I1["Test Response"] --> I2["Test favicon sha256"]
        I2 --> I3["Test web"]
        I3 -->|"report"| SERVER
    end
    ABC ----->|"report"| SERVER
```

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
