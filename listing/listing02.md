Что такое интерфейсы, как они применяются в Go?

Ответ:

Интерфейс - абстракция, которая описывает поведение других типов. Он представляет собой набор методов
без их реализации. Большой плюс интерфейсов заключается в том, что мы не зависим от конечной реализации.
Мы можем передавать различные аргументы и получать итоговый результат.
Интерфейсы используются в Go для реализации полиморфизма. Это позволяет создавать гибкие компоненты, 
реализация которых может быть заменена на другую без изменения кода.