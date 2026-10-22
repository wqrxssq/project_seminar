# Квартирник

Цифровой паспорт квартиры: документы, техника, гарантии, платежи и показания у арендодателя и арендатора в одном сервисе.

## One-pager

Презентационный лист лежит в [`one_pager/`](one_pager/):

- [`one-pager.html`](one_pager/one-pager.html) — исходная вёрстка, один вертикальный лист 210×478 мм
- [`one-pager.pdf`](one_pager/one-pager.pdf) — готовый PDF
- [`one-pager.png`](one_pager/one-pager.png) — превью

Открыть вёрстку в браузере:

```bash
open one_pager/one-pager.html
```

## Презентация

Статус сделанного — в [`docs/kt1.pptx`](docs/kt1.pptx): C4, роли микросервисов, выбор River и план работ.

## Backend

Черновик Go-бэкенда — в [`backend/`](backend/README.md). Сейчас сервисы отвечают только на `GET /ping`. Домены и границы описаны в [`docs/c4.md`](docs/c4.md).
