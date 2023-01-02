# grabber_go
 Есть в интернете сайт http://vpustotu.ru на котором любой желающий может анонимно высказаться о наболевшем. Все высказывания (в дальнейшем буду называть их “цитатами”) сначала попадают в модерацию (аналог “бездны” башорга), где посетители могут оценить полет мысли и проголосовать за цитату в стиле “Ого!” или “Ерунда!”. На странице модерации (http://vpustotu.ru/moderation/) нам показывают случайную цитату, ссылки голосования и ссылку “Еще”, которая ведет на эту же страницу. Пощелкайте, это все очень просто.

И вот возникла задача – срочно, под покровом темноты, загрузить себе полный дамп всех цитат на модерации для дальнейшего секретного исследования. Не будем оценивать житейскую ценность и степень идиотизма задачи, а рассмотрим её с технической точки зрения:

В разделе модерации нет прямых ссылок на определенную цитату, единственный способ получить новую цитату – обновить страницу (или перейти по ссылке “еще”, что одно и тоже). Причем вполне возможны повторы, что легко обнаруживается после пары минут агрессивного кликинга.

Таким образом нужна программа, которая:

Должна последовательно обновлять и парсить (разбирать) страницу, записывая цитату.
Должна уметь отбрасывать дубликаты.
Логично, что мы понятия не имеем все ли цитаты загружены, но об этом можно косвенно догадаться по большому количеству повторно полученных цитат подряд. Поэтому дополним:

Должна останавливаться не только по команде, но и по достижению определенного числа “повторов”, например 500!
Так как это, скорее всего, займет некоторое время: необходимо уметь продолжить “с места на котором остановились” после закрытия.
Ну и раз уж все-таки это надолго – пусть делает свое грязное дело в несколько потоков. Хорошо-бы в целых 4 потока (или даже 5!).
И отчитывается об успехах в консоль каждые, скажем, 10 секунд.
А все эти параметры пускай принимает из аргументов командной строки!
