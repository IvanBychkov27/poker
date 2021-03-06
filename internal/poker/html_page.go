package poker

const (
	pageBottom = `
<form action="/" method="GET">
	<button type="submit" name="resetButton" value="reset">Очистить</button>
<form>
<br>
<p><H5><I>IvanBychkov@mail.ru</I></H5>
<p><H6><I><span style="color: lightgray"> {count_visits} </span></I></H6>
</body></html>`
	pageTop = `
<!DOCTYPE HTML>
<html>
	<head>
		<link rel="icon" href="http://poker-iv.herokuapp.com/favicon.ico" type="image/x-icon">
		<meta charset="utf-8">
		<title>Poker combinations</title>
		<style>
        table {
            width: 1300px;
            /* Ширина таблицы */
            border: black 1px solid;
            /* Стиль рамки таблицы */
            border-spacing: 0;
            /* Расстояние между ячейками */
            border-collapse: collapse; /* Одинарная рамка */
            margin: 0px;
            /* Выравнивание таблицы по центру */
        }

        caption {
            text-align: center;
            /* Выравнивание заголовка по центру */
            caption-side: top
            /* Вывод заголовка над таблицей */
        }

        td, th {
            border: black 1px solid;
            /* Стиль рамки ячеек */
            padding: 10px;
            /* Отступ между границей и содержимым */
            text-align: center
            /* Выравнивание ячейки по центру */
        }

        tr {
            text-align: center;
            /* Горизонтальное выравнивание */
            vertical-align: top;
            /* Вертикальное выравнивание */
        }

        tr:nth-child(2n) {
            background: #e8e8e8
            /* Зебра */
        }

        thead, tfoot {
            background: #f5e0cd
            /* Цвет фона */
        }
    </style>

    <script src="https://www.google.com/jsapi"></script>
    <script>
        google.load("visualization", "1", {packages:["corechart"]});
        google.setOnLoadCallback(drawChart);
        function drawChart() {
            var data = google.visualization.arrayToDataTable({chart_comb});
            var options = {
                title: 'Комбинации',
                is3D: true,
                pieResidueSliceLabel: 'Остальное'
            };
            var chart = new google.visualization.PieChart(document.getElementById('air'));
            chart.draw(data, options);
        }
    </script>

	</head>
<body>
`

	form = `
<form action="/" method="GET">
<table>
        <caption><H2><B>Покерная вероятность победы</B></H2></caption>
        <thead>
        <tr>
            <th>Введите <B>две Ваши</B> карты:</th>
            <th>Ваши карты</th>
            <th>Ваша комбинация</th>
        </tr>
        </thead>
        <tbody>
        <tr>
            <td>
                <p>&emsp; &emsp; &nbsp; 2 &ensp; 3 &ensp; 4 &ensp; 5 &ensp; 6 &ensp; 7 &ensp; 8&ensp; 9&ensp; 10&ensp; V&ensp;
                    D &ensp; K &ensp; T</p>
                <p>Пик&nbsp;
                    <input type="checkbox" name="hand2p"/>
                    <input type="checkbox" name="hand3p"/>
                    <input type="checkbox" name="hand4p"/>
                    <input type="checkbox" name="hand5p"/>
                    <input type="checkbox" name="hand6p"/>
                    <input type="checkbox" name="hand7p"/>
                    <input type="checkbox" name="hand8p"/>
                    <input type="checkbox" name="hand9p"/>
                    <input type="checkbox" name="hand10p"/>
                    <input type="checkbox" name="hand11p"/>
                    <input type="checkbox" name="hand12p"/>
                    <input type="checkbox" name="hand13p"/>
                    <input type="checkbox" name="hand14p"/>
                </p>
                <p>Крес
                    <input type="checkbox" name="hand2k"/>
                    <input type="checkbox" name="hand3k"/>
                    <input type="checkbox" name="hand4k"/>
                    <input type="checkbox" name="hand5k"/>
                    <input type="checkbox" name="hand6k"/>
                    <input type="checkbox" name="hand7k"/>
                    <input type="checkbox" name="hand8k"/>
                    <input type="checkbox" name="hand9k"/>
                    <input type="checkbox" name="hand10k"/>
                    <input type="checkbox" name="hand11k"/>
                    <input type="checkbox" name="hand12k"/>
                    <input type="checkbox" name="hand13k"/>
                    <input type="checkbox" name="hand14k"/>
                </p>
                <p>Буби
                    <input type="checkbox" name="hand2b"/>
                    <input type="checkbox" name="hand3b"/>
                    <input type="checkbox" name="hand4b"/>
                    <input type="checkbox" name="hand5b"/>
                    <input type="checkbox" name="hand6b"/>
                    <input type="checkbox" name="hand7b"/>
                    <input type="checkbox" name="hand8b"/>
                    <input type="checkbox" name="hand9b"/>
                    <input type="checkbox" name="hand10b"/>
                    <input type="checkbox" name="hand11b"/>
                    <input type="checkbox" name="hand12b"/>
                    <input type="checkbox" name="hand13b"/>
                    <input type="checkbox" name="hand14b"/>
                </p>
                <p>Черв
                    <input type="checkbox" name="hand2ch"/>
                    <input type="checkbox" name="hand3ch"/>
                    <input type="checkbox" name="hand4ch"/>
                    <input type="checkbox" name="hand5ch"/>
                    <input type="checkbox" name="hand6ch"/>
                    <input type="checkbox" name="hand7ch"/>
                    <input type="checkbox" name="hand8ch"/>
                    <input type="checkbox" name="hand9ch"/>
                    <input type="checkbox" name="hand10ch"/>
                    <input type="checkbox" name="hand11ch"/>
                    <input type="checkbox" name="hand12ch"/>
                    <input type="checkbox" name="hand13ch"/>
                    <input type="checkbox" name="hand14ch"/>
                </p>
            </td>
            <td>{head_cards}</td>
            <td style="text-align: left">{head_victory}</td>
        </tr>
        <tr>
            <td>Введите <B>карты на столе</B>:</td>
            <td>Карты на столе</td>
            <td>Возможные варианты комбинаций</td>
        </tr>
        <tr>
            <td>
                <p>&emsp; &emsp; &nbsp; 2 &ensp; 3 &ensp; 4 &ensp; 5 &ensp; 6 &ensp; 7 &ensp; 8&ensp; 9&ensp; 10&ensp; V&ensp;
                    D &ensp; K &ensp; T</p>
                <p>Пик&nbsp;
                    <input type="checkbox" name="table2p"/>
                    <input type="checkbox" name="table3p"/>
                    <input type="checkbox" name="table4p"/>
                    <input type="checkbox" name="table5p"/>
                    <input type="checkbox" name="table6p"/>
                    <input type="checkbox" name="table7p"/>
                    <input type="checkbox" name="table8p"/>
                    <input type="checkbox" name="table9p"/>
                    <input type="checkbox" name="table10p"/>
                    <input type="checkbox" name="table11p"/>
                    <input type="checkbox" name="table12p"/>
                    <input type="checkbox" name="table13p"/>
                    <input type="checkbox" name="table14p"/>
                </p>
                <p>Крес
                    <input type="checkbox" name="table2k"/>
                    <input type="checkbox" name="table3k"/>
                    <input type="checkbox" name="table4k"/>
                    <input type="checkbox" name="table5k"/>
                    <input type="checkbox" name="table6k"/>
                    <input type="checkbox" name="table7k"/>
                    <input type="checkbox" name="table8k"/>
                    <input type="checkbox" name="table9k"/>
                    <input type="checkbox" name="table10k"/>
                    <input type="checkbox" name="table11k"/>
                    <input type="checkbox" name="table12k"/>
                    <input type="checkbox" name="table13k"/>
                    <input type="checkbox" name="table14k"/>
                </p>
                <p>Буби
                    <input type="checkbox" name="table2b"/>
                    <input type="checkbox" name="table3b"/>
                    <input type="checkbox" name="table4b"/>
                    <input type="checkbox" name="table5b"/>
                    <input type="checkbox" name="table6b"/>
                    <input type="checkbox" name="table7b"/>
                    <input type="checkbox" name="table8b"/>
                    <input type="checkbox" name="table9b"/>
                    <input type="checkbox" name="table10b"/>
                    <input type="checkbox" name="table11b"/>
                    <input type="checkbox" name="table12b"/>
                    <input type="checkbox" name="table13b"/>
                    <input type="checkbox" name="table14b"/>
                </p>
                <p>Черв
                    <input type="checkbox" name="table2ch"/>
                    <input type="checkbox" name="table3ch"/>
                    <input type="checkbox" name="table4ch"/>
                    <input type="checkbox" name="table5ch"/>
                    <input type="checkbox" name="table6ch"/>
                    <input type="checkbox" name="table7ch"/>
                    <input type="checkbox" name="table8ch"/>
                    <input type="checkbox" name="table9ch"/>
                    <input type="checkbox" name="table10ch"/>
                    <input type="checkbox" name="table11ch"/>
                    <input type="checkbox" name="table12ch"/>
                    <input type="checkbox" name="table13ch"/>
                    <input type="checkbox" name="table14ch"/>
                </p>

            </td>
            <td>{table_cards}</td>
            <td style="text-align: left">
				{table_victory}
				<div id="air" style="width: 350px; height: 200px;"></div>
			</td>
        </tr>

        <tr>
            <td>
                <details>
                    <summary><B>Вышедшие карты: </B></summary>
                    <p>&emsp; &emsp; &nbsp; 2 &ensp; 3 &ensp; 4 &ensp; 5 &ensp; 6 &ensp; 7 &ensp; 8&ensp; 9&ensp; 10&ensp;
                        V&ensp;
                        D &ensp; K &ensp; T</p>
                    <p>Пик&nbsp;
                        <input type="checkbox" name="out2p"/>
                        <input type="checkbox" name="out3p"/>
                        <input type="checkbox" name="out4p"/>
                        <input type="checkbox" name="out5p"/>
                        <input type="checkbox" name="out6p"/>
                        <input type="checkbox" name="out7p"/>
                        <input type="checkbox" name="out8p"/>
                        <input type="checkbox" name="out9p"/>
                        <input type="checkbox" name="out10p"/>
                        <input type="checkbox" name="out11p"/>
                        <input type="checkbox" name="out12p"/>
                        <input type="checkbox" name="out13p"/>
                        <input type="checkbox" name="out14p"/>
                    </p>
                    <p>Крес
                        <input type="checkbox" name="out2k"/>
                        <input type="checkbox" name="out3k"/>
                        <input type="checkbox" name="out4k"/>
                        <input type="checkbox" name="out5k"/>
                        <input type="checkbox" name="out6k"/>
                        <input type="checkbox" name="out7k"/>
                        <input type="checkbox" name="out8k"/>
                        <input type="checkbox" name="out9k"/>
                        <input type="checkbox" name="out10k"/>
                        <input type="checkbox" name="out11k"/>
                        <input type="checkbox" name="out12k"/>
                        <input type="checkbox" name="out13k"/>
                        <input type="checkbox" name="out14k"/>
                    </p>
                    <p>Буби
                        <input type="checkbox" name="out2b"/>
                        <input type="checkbox" name="out3b"/>
                        <input type="checkbox" name="out4b"/>
                        <input type="checkbox" name="out5b"/>
                        <input type="checkbox" name="out6b"/>
                        <input type="checkbox" name="out7b"/>
                        <input type="checkbox" name="out8b"/>
                        <input type="checkbox" name="out9b"/>
                        <input type="checkbox" name="out10b"/>
                        <input type="checkbox" name="out11b"/>
                        <input type="checkbox" name="out12b"/>
                        <input type="checkbox" name="out13b"/>
                        <input type="checkbox" name="out14b"/>
                    </p>
                    <p>Черв
                        <input type="checkbox" name="out2ch"/>
                        <input type="checkbox" name="out3ch"/>
                        <input type="checkbox" name="out4ch"/>
                        <input type="checkbox" name="out5ch"/>
                        <input type="checkbox" name="out6ch"/>
                        <input type="checkbox" name="out7ch"/>
                        <input type="checkbox" name="out8ch"/>
                        <input type="checkbox" name="out9ch"/>
                        <input type="checkbox" name="out10ch"/>
                        <input type="checkbox" name="out11ch"/>
                        <input type="checkbox" name="out12ch"/>
                        <input type="checkbox" name="out13ch"/>
                        <input type="checkbox" name="out14ch"/>
                    </p>

                </details>
            </td>

            <td>{released_cards}</td>
            <td> </td>
        </tr>

        </tbody>
</table>

<br>
<p>Количество соперников:
    <input type="number" name="nPlayers" value="1"  min="1" max="15" size="1" step="1">
</p>
<br>
	<input type="submit" value="Рассчитать вероятность">
<br>
<br>==============================<br><br>
</form>
`
)
