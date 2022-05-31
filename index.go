
package main
 
import (
    "fmt"
    
)



func start(){
    cups := 9
    money := 390
    water := 540
    maxWater := 5000
    maxMilk := 1000
    maxCoffe := 900
    maxCups := 50
    coffe := 120
    milk := 400 
    var selectedCommand string
    fmt.Println("Введите команд(buy, fill, take, stat, exit):")
    fmt.Scanf("%s\n", &selectedCommand)

    for selectedCommand != "exit" {
        switch selectedCommand {
            case "buy":
                var drinkType string
                fmt.Println("Что вы хотите купить \n - Espresso \n - Latte \n - Cappuccino")
                fmt.Scanf("%s\n", &drinkType)

                switch drinkType{
                    case "Espresso":
                        if  water < 250 || coffe < 16 || cups == 0{
                            fmt.Println("В аппарате недостаточное количества воды или кофе, или закончились стаканчики")
                        } else {
                            fmt.Printf("ВЫ выбрали напиток: %s\n", drinkType)
                            water = water - 250
                            coffe = coffe - 16
                            money = money + 60
                            cups = cups - 1                            
                        }

                    case "Latte": 
                        if  water < 300 || coffe < 20 || cups == 0 || milk < 76{
                            fmt.Println("В аппарате недостаточное количества воды или кофе или молока, или закончились стаканчики")
                        } else {
                            fmt.Printf("ВЫ выбрали напиток: %s\n", drinkType)
                            water = water - 300
                            coffe = coffe - 20
                            milk = milk - 76
                            money = money + 100
                            cups = cups - 1                            
                        }
                    case "Cappuccino":
                        if  water < 200 || coffe < 16 || cups == 0 || milk < 100{
                            fmt.Println("В аппарате недостаточное количества воды или кофе или молока, или закончились стаканчики")   
                        } else {
                            fmt.Printf("ВЫ выбрали напиток: %s\n", drinkType)
                            water = water - 200
                            coffe = coffe - 16
                            milk = milk - 100
                            money = money + 110
                            cups = cups - 1
                        }
                    default:
                        fmt.Printf("Выбранный напиток не доступен")
                }
            case "fill":
                var fillType string
                fmt.Printf("Что вы хотите пополнить: \n - Water \n - Milk \n - Coffe \n - Cups \n")
                fmt.Scanf("%s\n", &fillType)
                switch fillType{
                    case "Water":
                        var waterCount, a int
                        a = maxWater - water 
                        fmt.Println("Текущий остаток воды: \n", water)
                        fmt.Println("Максимальный объем: \n", maxWater)
                        fmt.Println("Вы можете добавить воды до: \n", a)
                        fmt.Println("Какой объем залить воды?")
                        fmt.Scanf("%d\n", &waterCount)
                        if waterCount > a {
                            fmt.Println("Вы пытаетесь залить воды больше объема бака")
                        } else {
                            fmt.Println("Вы залили: ", waterCount)
                            water = water + waterCount
                            fmt.Println("Текущий объем воды: ", water)
                        }
                    case "Milk":
                        var milkCount, b int
                        b = maxMilk - milk 
                        fmt.Println("Текущий остаток молока: \n", milk)
                        fmt.Println("Максимальный объем молока: \n", maxMilk)
                        fmt.Println("Вы можете добавить молока до: \n", b)
                        fmt.Println("Какой объем залить молока?")
                        fmt.Scanf("%d\n", &milkCount)
                        if milkCount > b {
                            fmt.Println("Вы пытаетесь залить молока больше объема бака")
                        } else {
                            fmt.Println("Вы залили: ", milkCount)
                            milk = milk + milkCount
                            fmt.Println("Текущий объем молока: ", milk)
                        }
                    case "Coffe":
                        var coffeCount, c int
                        c = maxCoffe - coffe 
                        fmt.Println("Текущий остаток зерен кофе: \n", coffe)
                        fmt.Println("Максимальный объем зерен кофе: \n", maxCoffe)
                        fmt.Println("Вы можете добавить зерен кофе до: \n", c)
                        fmt.Println("Сколько грамм кофе засыпать?")
                        fmt.Scanf("%d\n", &coffeCount)
                        if coffeCount > c {
                            fmt.Println("Вы пытаетесь добавить слишком много зерен кофе")
                        } else {
                            fmt.Println("Вы залили: ", coffeCount)
                            coffe = coffe + coffeCount
                            fmt.Println("Текущая масса кофе: ", coffe)
                        }
                    case "Cups":
                        var cupsCount, d int
                        d = maxCups - cups 
                        fmt.Println("Текущий остаток стаканчиков: \n", cups)
                        fmt.Println("Максимальное количество стаканчиков: \n", maxCups)
                        fmt.Println("Вы можете добавить стаканчиков до: \n", d)
                        fmt.Println("Сколько добавить стаканчиков?")
                        fmt.Scanf("%d\n", &cupsCount)
                        if cupsCount > d {
                            fmt.Println("Вы пытаетесь добавить слишком много стаканчиков")
                        } else {
                            fmt.Println("Вы добавили стаканчиков: ", cupsCount)
                            cups = cups + cupsCount
                            fmt.Println("Текущее количество стаканчиков: ", cups)
                        }

                }
            case "take":
                fmt.Println("Вы опусташили кассу, в кассе было:", money)
                money = 0
            case "stat":
                fmt.Printf("Текущий остаток: \n - стаканчиков: %+v шт\n - денег: %+v руб\n - воды: %+v мл\n - зерен: %+v гр\n - молока: %+v мл\n", cups, money, water, coffe, milk)
            case "exit":
                fmt.Printf("Выход из программы \n")
            default:
                fmt.Printf("Выбранной команды нет в списке \n")
            }
    fmt.Println("Введите команд(buy, fill, take, stat, exit):")
    fmt.Scanf("%s\n", &selectedCommand)
        }
    }



func main() {
 start();
}



    // var name string
    // fmt.Println("Как тебя зовут?")
    // fmt.Scanf("%s\n", &name)
 
    // var age int
    // fmt.Println("Сколько тебе лет?")
    // fmt.Scanf("%d\n", &age)
 
    // fmt.Printf("Привет, %s, твой возраст - %d\n", name, age)