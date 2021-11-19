# checkout-system
Golang Basics course Checkout System

### Задание

- task.docx (оригинальный текст задания) +
- Данные о продуктах считывать из json файла; 
- Путь к файлу указывать при помощи env переменных; 
- Не забывать об обработке ошибок; 
- Проект должен содержать Readme и Make files;  


### Примечания
- Так как в текущем состоянии это больше консольное приложение, поэтому решил сделать флаг -v для логов. Процесс выполнения пишется в fmt stdout

- Текущие допустимые скидки:
- - DiscountEveryMaxCountRule условие скидки - каждый (i == MaxCount) товар FromSKU будет пременена скидка ToPrice
- - DiscountMoreThanCountRule условие скидки - если количество FromSKU больше MoreThanCount - применяется цена ToPrice ко всем ToSKU
- - DiscountOneToOneRule условие скидки - на каждый один товар FromSKU - применяется цена ToPrice к одному ToSKU

- Чтобы создать новое условие для скидок нужно реализовать интерфейс DiscountRuler и добавить объект в PricingRules


- Наверное было бы классно сделать что-то вроде такого:
- - rules.From("ipd").Every(4).ToPrice(10000)
- - rules.From("ipd").To("mbp").If(func() bool).ToPrice(500)
- Но у меня явно не хватает опыта и знаний на это