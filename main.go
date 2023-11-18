package main

import (
	"fmt"
	"math"
	"os"
)

type ConversionRate struct {
	fromCurrency string
	toCurrency   string
	rate         float64
}

type Fraccion struct {
	numerador   int
	denominador int	
}

var conversionRates = []ConversionRate{
	{"USD", "EUR", 0.85},
	{"EUR", "USD", 1.18},
	{"USD", "JPY", 110.75},
	{"JPY", "USD", 0.009},
}

func findConversionRate(fromCurrency, toCurrency string) (float64, bool) {
	for _, rate := range conversionRates {
		if rate.fromCurrency == fromCurrency && rate.toCurrency == toCurrency {
			return rate.rate, true
		}
	}
	return 0.0, false
}

func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	if fromCurrency == toCurrency {
		return amount, nil
	}

	rate, found := findConversionRate(fromCurrency, toCurrency)
	if !found {
		return 0.0, fmt.Errorf("conversion rate not found for %s to %s", fromCurrency, toCurrency)
	}

	convertedAmount := amount * rate
	return convertedAmount, nil
}

func maximoComunDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mcm(a, b int) int {
	return a * b / maximoComunDivisor(a, b)
}

func minimoComunMultiplo(a, b int) {
	if a == 0 || b == 0 || a == 0 && b == 0 {
		os.Exit(1)
	}

	resultado := mcm(a, b)
	fmt.Printf("El minimo común múltiplo de %d y %d es %d", a, b, resultado)
}

func segundoGrado(a, b, c float64) {
	discriminador := b*b - 4*a*c
	if discriminador < 0 {
		fmt.Printf("La ecuación ((%.2fx^2 + %.2fx + %.2f) no tiene soluciones posibles", a, b, c)
	} else if discriminador == 0 {
		x := -b / (2 * a)
		fmt.Printf("La única solución de la ecuación (%.2fx^2 + %.2fx + %.2f) es %.2f", a, b, c, x)
	} else {
		x1 := (-b + math.Sqrt(float64(discriminador))) / (2 * a)
		x2 := (-b - math.Sqrt(float64(discriminador))) / (2 * a)
		fmt.Printf("Las soluciones de la ecuación (%.2fx^2 + %.2fx + %.2f) son: x1 = %.2f y x2 = %.2f", a, b, c, float64(x1), float64(x2))
	}
}

func pitagoras(a, b int) {
	b2 := a * a
	c2 := b * b
	resultado := math.Sqrt(float64(b2 + c2))
	fmt.Printf("La hipotenusa del triangulo con los catetos de %.2f y %.2f es %.2f", float64(a), float64(b), resultado)
}

func notacion(a, b int) {
	notacion := float64(a) * (math.Pow(10, float64(b)))
	fmt.Printf("La notación científica de %.2f elevado a %.2f es %.2f", float64(a), float64(b), notacion)
}

func newtonsKilogramos(input int) {
	var a int

	switch input {
	case 1:
		fmt.Print("Ingresa la cantidad de newtons: ")
		fmt.Scan(&a)

		resultado := float64(a) / 9.8
		fmt.Printf("%.2f newtons son %.2f kilogramos", float64(a), resultado)
	case 2:
		fmt.Print("Ingresa la cantidad de kilogramos: ")
		fmt.Scan(&a)

		resultado := float64(a) * 9.8
		fmt.Printf("%.2f kilogramos son %.2f newtons", float64(a), resultado)
	}
}

func areaPerimetroInput(input int) {
	var input1 int

	switch input {
	case 1:
		fmt.Println("[1] Calcular el área de un triangulo rectángulo")
		fmt.Println("[2] Calcular el área de un triangulo")
		fmt.Println("[3] Calcular el área de un rectangulo / cuadrado")
		fmt.Println("[4] Calcular el área de un circulo")
		fmt.Println("[5] Calcular el área de un rombo")
		fmt.Println("[6] Calcular el área de un paralelogramo")
		fmt.Println("[7] Calcular el área de un trapecio")
		fmt.Println("[8] Calcular el área de un poligono regular")
		fmt.Print("Opción >> ")
		fmt.Scan(&input1)

		calcularArea(input1)
	case 2:
		fmt.Println("[1] Calcular el perímetro de un triangulo rectángulo")
		fmt.Println("[2] Calcular el perímetro de un rectangulo / cuadrado")
		fmt.Println("[3] Calcular el perímetro de un circulo")
		fmt.Println("[4] Calcular el perímetro de un rombo")
		fmt.Println("[5] Calcular el perímetro de un paralelogramo")
		fmt.Println("[6] Calcular el perímetro de un trapecio")
		fmt.Println("[7] Calcular el perímetro de un poligono regular")
		fmt.Print("Opcion >> ")
		fmt.Scan(&input1)

		calcularPerimetro(input1)
	}
}

func calcularPerimetro(input int) {
	switch input {
	case 1:
		var a, b, c float64

		fmt.Print("Ingresa los valores de los lados del triangulo >> ")
		fmt.Scan(&a, &b, &c)

		resultado := a + b + c
		fmt.Printf("El perímetro del triangulo es %.2f", resultado)
	case 2:
		var a, b float64

		fmt.Print("Ingresa los valores de la base y la altura >> ")
		fmt.Scan(&a, &b)

		resultado := 2*a + 2*b
		fmt.Printf("El perimetro del cuadrado es %.2f", resultado)
	case 3:
		var r float64

		fmt.Print("Ingresa el valor del radio del circulo >> ")
		fmt.Scan(&r)

		resultado := 2 * 3.14 * r
		fmt.Printf("El perimetro del circulo es de %.2f", resultado)
	case 4:
		var l float64

		fmt.Print("Ingresa el valor del lado del rombo >> ")
		fmt.Scan(&l)

		resultado := 4 * l
		fmt.Printf("El perimetro del rombo es de %.2f", resultado)
	case 5:
		var b, c float64

		fmt.Print("Ingresa el valor de la altura y la diagonal del paralelogramo >> ")
		fmt.Scan(&b, &c)

		resultado := 2 * (b + c)
		fmt.Printf("El perimetro del paralelogramo es de %.2f", resultado)
	case 6:
		var B, c, b, d float64

		fmt.Print("Ingresa el valor de la base mayor, base menor, diagonal 1 y diagonal 2 >> ")
		fmt.Scan(&B, &b, &c, &d)

		resultado := B + c + b + d
		fmt.Printf("El perimetro del trapecio es de %.2f", resultado)
	case 7:
		var n, l float64

		fmt.Printf("Ingresa el valor del lado del poligino y el numero de lados >> ")
		fmt.Scan(&l, &n)

		resultado := n * l
		fmt.Printf("El perimetro del poligono de %d lados es de %.2f", int(n), resultado)
	}
}

func calcularArea(input int) {
	switch input {
	case 1:
		var base, altura float64

		fmt.Print("Ingresa los valores de la base y la altura >> ")
		fmt.Scan(&base, &altura)

		resultado := (base * altura) / 2

		fmt.Printf("El área del triangulo con base %.2f y altura %.2f es %.2f", base, altura, resultado)
	case 2:
		var a, b, c float64

		fmt.Print("Ingresa los valores de los lados del triangulo >> ")
		fmt.Scan(&a, &b, &c)

		s := (a + b + c) / 2

		area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
		fmt.Printf("El área del triangulo cuyos lados son %.2f, %.2f y %.2f es %.2f", a, b, c, area)
	case 3:
		var a, b float64

		fmt.Print("Ingresa el valor de los dos lados >> ")
		fmt.Scan(&a, &b)

		resultado := a * b
		if a != b {
			fmt.Printf("El área del rectangulo cuyo lado es %.2f y %.2f es %.2f", a, b, resultado)
		} else {
			fmt.Printf("El área del cuadrado cuyo lado es %.2f es %.2f", a, resultado)
		}
	case 4:
		var r float64

		fmt.Print("Ingresa el valor del radio del circulo >> ")
		fmt.Scan(&r)

		resultado := 3.14 * (r * r)
		fmt.Printf("El área del circulo cuyo radio es de %.2f es %.2f", r, resultado)
	case 5:
		var D, d float64

		fmt.Print("Ingresa el valor de la diagonal mayor y diagonal menor >> ")
		fmt.Scan(&D, &d)

		resultado := (D * d) / 2
		fmt.Printf("El área de un rombo cuyas diagonales son %.2f y %.2f es de %.2f", D, d, resultado)
	case 6:
		var a, b float64

		fmt.Print("Ingresa el valor de la base y la altura >> ")
		fmt.Scan(&a, &b)

		resultado := a * b
		fmt.Printf("El área de un pararelogramo cuya base es %.2f y altura es %.2f es de %.2f", a, b, resultado)
	case 7:
		var B, b, a float64

		fmt.Print("Ingresa el valor de la base mayor, base menor y altura del trapecio >> ")
		fmt.Scan(&B, &b, &a)

		resultado := ((B + b) / 2) * a
		fmt.Printf("El área del trapecio cuya base mayor es de %.2f, base menor %.2f y altura %.2f es de %.2f", B, b, a, resultado)
	case 8:
		var l, n, a float64

		fmt.Print("Ingresa el valor del lado, numero de lados, y apotema del poligono >> ")
		fmt.Scan(&l, &n, &a)

		p := n * l
		resultado := (0.5 * p) * a
		fmt.Printf("El área del poligono regular de %.2f lados, cuyo apotema es de %.2f y lado de %.2f es de %.2f", n, a, l, resultado)

	}
}

func ecuacionLineal(a, b float64) {
	if a != 0 {
		solucion := -b / a
		fmt.Printf("La solución de la ecuación (%.2fx + %.2f) es %.2f\n", a, b, solucion)
	} else {
		if b == 0 {
			fmt.Println("La ecuación tiene infinitas soluciones")
		} else {
			fmt.Println("La ecuación no tiene soluciones")
		}
	}
}

func coeficienteBinomial(n, k int) int {
	if k > n-k {
		k = n - k
	}

	numerador := 1
	denominador := 1

	for i := 0; i < k; i++ {
		numerador *= n - i
		denominador *= i + 1
	}

	return numerador / denominador
}

func binomioDeNewton(a, b, n int) {
	for k := 0; k <= n; k++ {
		coefBinomial := coeficienteBinomial(n, k)
		termA := int(math.Pow(float64(a), float64(n-k)))
		termB := int(math.Pow(float64(b), float64(k)))
		resultado := coefBinomial * termA * termB

		fmt.Printf("(%d + %d)^%d = %d\n", a, b, n, resultado)
	}
}

func sumarFracciones(frac1, frac2 Fraccion) Fraccion {
	suma := Fraccion{
		numerador:   frac1.numerador*frac2.denominador + frac2.numerador*frac1.denominador,
		denominador: frac1.denominador * frac2.denominador,
	}
	return simplificarFraccion(suma)
}

func restarFracciones(frac1, frac2 Fraccion) Fraccion {
	resta := Fraccion{
		numerador:   frac1.numerador*frac2.denominador - frac2.numerador*frac1.denominador,
		denominador: frac1.denominador * frac2.denominador,
	}
	return simplificarFraccion(resta)
}

func multiplicarFracciones(frac1, frac2 Fraccion) Fraccion {
	multiplicacion := Fraccion{
		numerador:   frac1.numerador * frac2.numerador,
		denominador: frac1.denominador * frac2.denominador,
	}
	return simplificarFraccion(multiplicacion)
}

func dividirFracciones(frac1, frac2 Fraccion) Fraccion {
	division := Fraccion{
		numerador:   frac1.numerador * frac2.denominador,
		denominador: frac1.denominador * frac2.numerador,
	}
	return simplificarFraccion(division)
}

func simplificarFraccion(frac Fraccion) Fraccion {
	gcd := maximoComunDivisor(frac.numerador, frac.denominador)
	frac.numerador /= gcd
	frac.denominador /= gcd
	return frac
}

func options() {
	fmt.Println("CALCULADORA EN GO")
	fmt.Println("======================")
	fmt.Println("[1] Calcula el minimo común múltiplo")
	fmt.Println("[2] Resolver una ecuación de segundo grado (ax^2 + bx + c = 0)")
	fmt.Println("[3] Calcula la hipotenusa de un triangulo rectangulo (Pitagoras)")
	fmt.Println("[4] Calcula la notación científica")
	fmt.Println("[5] Convierte Newtons a Kilogramos (Y viceversa)")
	fmt.Println("[6] Calcular área / perímetro")
	fmt.Println("[7] Calcular operaciones aritmeticas básicas")
	fmt.Println("[8] Calcular potencias")
	fmt.Println("[9] Calcular logaritmos")
	fmt.Println("[10] Resolver ecuaciones lineales (ax + b = 0)")
	fmt.Println("[11] Calcular factoriales")
	fmt.Println("[12] Calcular binomio de Newton")
	fmt.Println("[13] Conversión de grados (Celsius, Fahrenheit y Kelvin)")
	fmt.Println("[14] Cálculo del área y volumen de figuras tridimensionales")
	fmt.Println("[15] Conversión de unidades")
	fmt.Println("[16] Calculadora trigonométrica")
	fmt.Println("[17] Graficar una función matemática [ROTO]")
	fmt.Println("[18] Calcular raíz cuadrada")
	fmt.Println("[19] Calcular raíz cúbica")
	fmt.Println("[20] Calcular otra raíz")
	fmt.Println("[21] Calcular un número primo")
	fmt.Println("[22] Calcular la aceleración")
	fmt.Println("[23] Sumar fracciones")
	fmt.Println("[24] Restar fracciones")
	fmt.Println("[25] Multiplicar fracciones")
	fmt.Println("[26] Divir fracciones")
	fmt.Println("[27] Calcular la masa")
	fmt.Println("[28] Calculadora de álgebra booleana")
	fmt.Println("[29] Calcular la fuerza")
	fmt.Println("[30] Conversión de monedas")
	fmt.Println("[31] Calcular impuestos")
	fmt.Println("[32] Calcular notas")
	fmt.Println("[33] Calcular energía")
	fmt.Println("[34] Calcular crecimiento exponencial")
	fmt.Println("[35] Calcular prestamos")
	fmt.Print("Opción >> ")
}

func crecimientoExponencial(valorInicial, tasaCrecimiento float64, periodos int) float64 {
	return valorInicial * math.Pow(1+tasaCrecimiento, float64(periodos))
}

func calcularRaiz(n float64, indice int) float64 {
	return math.Pow(n, 1.0/float64(indice))
}

func conversionGrados(input int) {
	var a float64

	switch input {
	case 1:
		fmt.Print("Ingresa el valor en Celsius >> ")
		fmt.Scan(&a)

		resultado := a*1.8 + 32.00
		fmt.Printf("%.2f grados Celsius a Fahrenheit son %.2f", a, resultado)
	case 2:
		fmt.Print("Ingresa el valor en Celsius >> ")
		fmt.Scan(&a)

		resultado := a + 273.15
		fmt.Printf("%.2f grados Celsius son %.2f Kelvin", a, resultado)
	case 3:
		fmt.Print("Ingresa el valor en Fahrenheit >> ")
		fmt.Scan(&a)

		resultado := float64(a-32) / 1.8
		fmt.Printf("%.2f Grados Fahrenheit son %.2f Celsius", a, resultado)
	case 4:
		fmt.Print("Ingresa el valor en Fahrenheit >> ")
		fmt.Scan(&a)

		resultado := float64(a-32)*5.0/9.0 + 273.15
		fmt.Printf("%.2f Fahrenheit son %.2f Kelvin", a, resultado)
	case 5:
		fmt.Print("Ingresa el valor en Celsius >> ")
		fmt.Scan(&a)

		resultado := a - 273.15
		fmt.Printf("%.2f Kelvin son %.2f Celsius", a, resultado)
	case 6:
		fmt.Print("Ingresa el valor en Celsius >> ")
		fmt.Scan(&a)

		resultado := (a-273.15)*9/5 + 32
		fmt.Printf("%.2f Kelvin son %.2f Fahrenheit", a, resultado)
	}
}

func aritmetica(input int) {
	var a, b int

	switch input {
	case 1:
		fmt.Print("Ingresa los dos números a sumar >> ")
		fmt.Scan(&a, &b)

		resultado := a + b
		fmt.Printf("La suma entre %d y %d es %d", a, b, resultado)
	case 2:
		fmt.Print("Ingresa los dos números a restar: ")
		fmt.Scan(&a, &b)
		resultado := a - b
		fmt.Printf("La resta entre %d y %d es %d", a, b, resultado)
	case 3:
		fmt.Print("Ingresa los dos números a multiplicar: ")
		fmt.Scan(&a, &b)
		resultado := a * b
		fmt.Printf("La multiplicación entre %d y %d es %d", a, b, resultado)
	case 4:
		fmt.Print("Ingresa los dos números a dividir: ")
		fmt.Scan(&a, &b)
		resultado := a / b
		fmt.Printf("La división entre %d y %d es %d", a, b, resultado)
	default:
		fmt.Println("Opción inválida")
		os.Exit(1)
	}
}

func areaCubo(lado float64) float64 {
	return 6 * lado * lado
}

func volumenCubo(lado float64) float64 {
	return lado * lado * lado
}

func areaEsfera(radio float64) float64 {
	return 4 * math.Pi * radio * radio
}

func volumenEsfera(radio float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(radio, 3)
}

func calcularFactorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * calcularFactorial(n-1)
}

func calcularPotencias(base, exponente int) {
	resultado := math.Pow(float64(base), float64(exponente))

	fmt.Printf("El resultado de la potencia es %.2f", resultado)
}

func calcularLogaritmos(input int) {
	var numero int
	fmt.Print("Ingresa el número >> ")
	fmt.Scan(&numero)

	switch input {
	case 1:
		resultado := math.Log(float64(numero))
		fmt.Printf("El logaritmo natural de %d es %.2f", numero, resultado)
	case 2:
		resultado := math.Log10(float64(numero))
		fmt.Printf("El logaritmo de base 10 de %d es %.2f", numero, resultado)
	}
}

func calcularAreaVolumen(input int) {
	var resultado float64

	switch input {
	case 1:
		var lado float64
		fmt.Print("Ingresa la longitud del lado del cubo: ")
		fmt.Scan(&lado)
		resultado = areaCubo(lado)
		fmt.Printf("El área del cubo es: %.2f\n", resultado)
	case 2:
		var lado float64
		fmt.Print("Ingresa la longitud del lado del cubo: ")
		fmt.Scan(&lado)
		resultado = volumenCubo(lado)
		fmt.Printf("El volumen del cubo es: %.2f\n", resultado)
	case 3:
		var radio float64
		fmt.Print("Ingresa el radio de la esfera: ")
		fmt.Scan(&radio)
		resultado = areaEsfera(radio)
		fmt.Printf("El área de la esfera es: %.2f\n", resultado)
	case 4:
		var radio float64
		fmt.Print("Ingresa el radio de la esfera: ")
		fmt.Scan(&radio)
		resultado = volumenEsfera(radio)
		fmt.Printf("El volumen de la esfera es: %.2f\n", resultado)
	default:
		fmt.Println("Opción Inválida")
	}
}

func metrosAPies(metros float64) float64 {
	return metros * 3.28084
}

func piesAMetros(pies float64) float64 {
	return pies / 3.28084
}

func gramosALibras(gramos float64) float64 {
	return gramos * 0.00220462
}

func librasAGramos(libras float64) float64 {
	return libras / 0.00220462
}

func conversionUnidades(input int) {
	var resultado float64

	switch input {
	case 1:
		var metros float64
		fmt.Print("Ingresa la cantidad en metros: ")
		fmt.Scan(&metros)
		resultado = metrosAPies(metros)
		fmt.Printf("%.2f metros son %.2f pies\n", metros, resultado)
	case 2:
		var pies float64
		fmt.Print("Ingresa la cantidad en pies: ")
		fmt.Scan(&pies)
		resultado = piesAMetros(pies)
		fmt.Printf("%.2f pies son %.2f metros\n", pies, resultado)
	case 3:
		var gramos float64
		fmt.Print("Ingresa la cantidad en gramos: ")
		fmt.Scan(&gramos)
		resultado = gramosALibras(gramos)
		fmt.Printf("%.2f gramos son %.2f libras\n", gramos, resultado)
	case 4:
		var libras float64
		fmt.Print("Ingresa la cantidad en libras: ")
		fmt.Scan(&libras)
		resultado = librasAGramos(libras)
		fmt.Printf("%.2f libras son %.2f gramos\n", libras, resultado)
	default:
		fmt.Println("Opción Inválida")
	}
}

func sin(angulo float64) float64 {
	return math.Sin(angulo)
}

func cos(angulo float64) float64 {
	return math.Cos(angulo)
}

func tan(angulo float64) float64 {
	return math.Tan(angulo)
}

func calcularTrigonometrica(input int) {
	var angulo float64

	fmt.Print("Ingresa el valor del ángulo en grados >> ")
	fmt.Scan(&angulo)

	switch input {
	case 1:
		resultado := sin(angulo * math.Pi / 180)
		fmt.Printf("El seno de %.2f grados es %.2f\n", angulo, resultado)
	case 2:
		resultado := cos(angulo * math.Pi / 180)
		fmt.Printf("El coseno de %.2f grados es %.2f\n", angulo, resultado)
	case 3:
		resultado := tan(angulo * math.Pi / 180)
		fmt.Printf("La tangente de %.2f grados es %.2f\n", angulo, resultado)
	default:
		fmt.Println("Opción Inválida")
	}
}

func and(a, b bool) bool {
	return a && b
}

func or(a, b bool) bool {
	return a || b
}

func not(a bool) bool {
	return !a
}

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func calculateTax(income float64, taxRate float64) float64 {
	taxAmount := income * (taxRate / 100)
	return taxAmount
}

func esPrimo(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func calculateAverageGrade(grades []float64) float64 {
	total := 0.0
	for _, grade := range grades {
		total += grade
	}
	average := total / float64(len(grades))
	return average
}

func calcularAceleracion(vinicial int, vfinal int, tiempo int) int {
	aceleracion := (vinicial - vfinal) / tiempo
	return aceleracion
}

func calcularEnergia(input int) {
	var energia float64
	var velocidad, masa, altura float64

	switch input {
	case 1:
		fmt.Print("Ingresa la velocidad >> ")
		fmt.Scan(&velocidad)
		fmt.Print("Ingresa la masa >> ")
		fmt.Scan(&masa)

		energia = 0.5 * masa * (velocidad * velocidad)
	case 2:
		fmt.Print("Ingresa la masa >> ")
		fmt.Scan(&masa)
		fmt.Print("Ingresa la velocidad >> ")
		fmt.Scan(&velocidad)

		energia = masa * altura * 9.8
	case 3:
		var opcion int
		var ec, ep float64

		fmt.Println("[1] Calcular Em con Ep y Ec")
		fmt.Println("[2] Calcular Em desde 0")
		fmt.Print("Opción >> ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Print("Ingresa Ep >> ")
			fmt.Scan(&ep)
			fmt.Print("Ingresa Ec >> ")
			fmt.Scan(&ec)
		case 2:
			fmt.Print("Ingresa la altura >> ")
			fmt.Scan(&altura)
			fmt.Print("Ingresa la velocidad >> ")
			fmt.Scan(&velocidad)
			fmt.Print("Ingresa la masa >> ")
			fmt.Scan(&masa)

			ec = 0.5 * masa * (velocidad * velocidad)
			ep = masa * 9.8 * altura
		}

		energia = ec + ep
	}

	fmt.Printf("La energia es de %.2f", energia)
}

func calcularPagoMensual(cantidadPrestamo, interesAnual float64, años int) float64 {
	interesMensual := interesAnual / 12 / 100
	numeroPagos := float64(años * 12)
	pago := (cantidadPrestamo * interesAnual) / (1 - math.Pow(1+interesMensual, -numeroPagos))
	return pago
}

func main() {
	var opcion int
	var num1, num2 int
	var input int

	options()
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		fmt.Print("Ingresa los dos números para calcular el MCM: ")
		fmt.Scan(&num1, &num2)

		minimoComunMultiplo(num1, num2)
	case 2:
		fmt.Print("Ingresa los 3 coeficientes de la ecuación (ax + bx + c = 0): ")
		var a, b, c float64
		fmt.Scan(&a, &b, &c)

		segundoGrado(a, b, c)
	case 3:
		fmt.Print("Ingresa la longitud de los dos catetos: ")
		fmt.Scan(&num1, &num2)

		pitagoras(num1, num2)
	case 4:
		fmt.Print("Ingresa el número y el exponente: ")
		fmt.Scan(&num1, &num2)

		notacion(num1, num2)
	case 5:
		fmt.Println("[1] Convertir de Newtons a Kilogramo")
		fmt.Println("[2] Convertir de Kilogramo a Newtons")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		newtonsKilogramos(input)
	case 6:
		fmt.Println("[1] Calcular área")
		fmt.Println("[2] Calcular perímetro")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		areaPerimetroInput(input)
	case 7:
		fmt.Println("[1] Sumar 2 valores")
		fmt.Println("[2] Restar 2 valores")
		fmt.Println("[3] Multiplicar 2 valores")
		fmt.Println("[4] Dividir 2 valores")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		aritmetica(input)
	case 8:
		var b, e int
		fmt.Print("Ingresa la base de la potencia >> ")
		fmt.Scan(&b)
		fmt.Print("Ingresa el exponente de la potencia >> ")
		fmt.Scan(&e)

		calcularPotencias(b, e)
	case 9:
		fmt.Println("[1] Calcular logaritmo natural")
		fmt.Println("[2] Calcular logaritmo de base 10")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		calcularLogaritmos(input)
	case 10:
		var a, b float64

		fmt.Print("Ingresa el coeficiente 'a' >> ")
		fmt.Scan(&a)

		fmt.Print("Ingresa el término independiente 'b' >> ")
		fmt.Scan(&b)

		ecuacionLineal(a, b)
	case 11:
		var num int
		fmt.Print("Ingresa el número para calcular el factorial >> ")
		fmt.Scan(&num)

		if num < 0 {
			fmt.Println("El número no puede ser negativo")
		} else {
			resultado := calcularFactorial(num)
			fmt.Printf("El factorial de %d es %d\n", num, resultado)
		}
	case 12:
		var n int

		fmt.Print("Ingresa el valor de 'a' para el binomio (a + b)^n >> ")
		fmt.Scan(&num1)
		fmt.Print("Ingresa el valor de 'b' para el binomio (a + b)^n >> ")
		fmt.Scan(&num2)
		fmt.Print("Ingresa el valor de 'n' para el binomi (a + b)^n >> ")
		fmt.Scan(&n)

		binomioDeNewton(num1, num2, n)
	case 13:
		fmt.Println("[1] Convertir de Celsius a Fahrenheit")
		fmt.Println("[2] Convertir de Celsius a Kelvin")
		fmt.Println("[3] Convertir de Fahrenheit a Celsius")
		fmt.Println("[4] Convertir de Fahrenheit a Kelvin")
		fmt.Println("[5] Convertir de Kelvin a Celsius")
		fmt.Println("[6] Convertir de Kelvin a Fahrenheit")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		conversionGrados(input)
	case 14:
		fmt.Println("[1] Calcular el área de un cubo")
		fmt.Println("[2] Calcular el volumen de un cubo")
		fmt.Println("[3] Calcular el área de una esfera")
		fmt.Println("[4] Calcular el volumen de una esfera")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		calcularAreaVolumen(input)
	case 15:
		fmt.Println("[1] Convertir de metros a pies")
		fmt.Println("[2] Convertir de pies a metros")
		fmt.Println("[3] Convertir de gramos a libras")
		fmt.Println("[4] Convertir de libras a gramos")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)

		conversionUnidades(input)
	case 16:
		fmt.Println("[1] Calcular el seno de un ángulo")
		fmt.Println("[2] Calcular el coseno de un ángulo")
		fmt.Println("[3] Calcular la tangente de un ángulo")
		fmt.Print("Opción >> ")
		fmt.Scan(&input)
		calcularTrigonometrica(input)
	case 18:
		fmt.Print("Ingresa el número para calcular su raíz cuadrada: ")
		fmt.Scan(&num1)
		resultado := calcularRaiz(float64(num1), 2)
		fmt.Printf("La raíz cuadrada de %.2f es %.2f\n", float64(num1), resultado)
	case 19:
		fmt.Print("Ingresa el número para calcular su raíz cúbica: ")
		fmt.Scan(&num1)
		resultado := calcularRaiz(float64(num1), 3)
		fmt.Printf("La raíz cúbica de %.2f es %.2f\n", float64(num1), resultado)
	case 20:
		var num2 int
		fmt.Print("Ingresa el número para calcular su raíz: ")
		fmt.Scan(&num1)
		fmt.Print("Ingresa el índice de la raíz: ")
		fmt.Scan(&num2)
		resultado := calcularRaiz(float64(num1), num2)
		fmt.Printf("La raíz de índice %d de %.2f es %.2f\n", num2, float64(num1), resultado)
	case 21:
		fmt.Print("Ingresa el número para verificar si es primo: ")
		fmt.Scan(&num1)
		if esPrimo(num1) {
			fmt.Printf("%d es un número primo\n", num1)
		} else {
			fmt.Printf("%d no es un número primo\n", num1)
		}
	case 22:
		var vinicial int
		var vfinal int
		var tiempo int
		fmt.Println("Ingresa la velocidad inicial >> ")
		fmt.Scan(&vinicial)
		fmt.Println("Ingresa la velocidad final >> ")
		fmt.Scan(&vfinal)
		fmt.Println("Ingresa el tiempo en segundos >> ")
		fmt.Scan(&tiempo)
		aceleracion := calcularAceleracion(vinicial, vfinal, tiempo)
		fmt.Printf("La aceleración en segundos es: %d m/s^2 \n", aceleracion)
	case 23:
		var frac1, frac2 Fraccion
		fmt.Println("Ingresa la primera fracción:")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac1.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac1.denominador)

		fmt.Println("Ingresa la segunda fracción:")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac2.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac2.denominador)

		suma := sumarFracciones(frac1, frac2)
		fmt.Printf("La suma de las fracciones es: %d/%d\n", suma.numerador, suma.denominador)
	case 24:
		var frac1, frac2 Fraccion
		fmt.Println("Ingresa la primera fracción >> ")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac1.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac1.denominador)

		fmt.Println("Ingresa la segunda fracción:")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac2.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac2.denominador)

		resta := restarFracciones(frac1, frac2)
		fmt.Printf("La resta de las fracciones es: %d/%d\n", resta.numerador, resta.denominador)
	case 25:
		var frac1, frac2 Fraccion
		fmt.Println("Ingresa la primera fracción >> ")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac1.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac1.denominador)

		fmt.Println("Ingresa la segunda fracción:")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac2.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac2.denominador)

		multiplicacion := multiplicarFracciones(frac1, frac2)
		fmt.Printf("La multiplicación de las fracciones es: %d/%d\n", multiplicacion.numerador, multiplicacion.denominador)
	case 26:
		var frac1, frac2 Fraccion
		fmt.Println("Ingresa la primera fracción >> ")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac1.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac1.denominador)

		fmt.Println("Ingresa la segunda fracción:")
		fmt.Print("Numerador >> ")
		fmt.Scan(&frac2.numerador)
		fmt.Print("Denominador >> ")
		fmt.Scan(&frac2.denominador)

		division := dividirFracciones(frac1, frac2)
		fmt.Printf("La división de las fracciones es: %d/%d\n", division.numerador, division.denominador)
	case 27:
		var newtons int
		fmt.Println("Introduce la fuerza en Newtons >> ")
		fmt.Scan(&newtons)
		var masa float64 = float64(newtons) / 9.81
		fmt.Printf("La masa es de: %fkg\n", masa)
	case 28:
		var a, b bool
		fmt.Print("Ingresa el valor de a (true o false) >> ")
		fmt.Scan(&a)
		fmt.Print("Ingresa el valor de b (true o false) >> ")
		fmt.Scan(&b)

		fmt.Printf("a AND b = %v\n", and(a, b))
		fmt.Printf("a OR b = %v\n", or(a, b))
		fmt.Printf("NOT a = %v\n", not(a))
		fmt.Printf("a XOR B = %v\n", xor(a, b))
	case 29:
		var masa float64
		fmt.Println("Ingresa la masa en Kilogramos >> ")
		fmt.Scan(&masa)
		var fuerza float64 = masa * 9.81
		fmt.Printf("La fuerza son: %f Newtons\n", fuerza)
	case 30:
		var amount float64
		var fromCurrency, toCurrency string

		fmt.Print("Ingresa la cantidad: ")
		fmt.Scan(&amount)
		fmt.Print("Ingresa la moneda de origen (ejemplo: USD, EUR, JPY): ")
		fmt.Scan(&fromCurrency)
		fmt.Print("Ingresa la moneda de destino (ejemplo: USD, EUR, JPY): ")
		fmt.Scan(&toCurrency)

		convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%.2f %s equivale a %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
	case 31:
		var income float64
		var taxRate float64

		fmt.Print("Ingresa el monto del ingreso: ")
		fmt.Scan(&income)
		fmt.Print("Ingresa la tasa impositiva (%): ")
		fmt.Scan(&taxRate)

		taxAmount := calculateTax(income, taxRate)
		fmt.Printf("El monto de impuestos a pagar es: %.2f\n", taxAmount)
	case 32:
		var numGrades int
		fmt.Print("Ingrese el número de notas: ")
		fmt.Scan(&numGrades)

		grades := make([]float64, numGrades)
		for i := 0; i < numGrades; i++ {
			fmt.Printf("Ingrese la nota %d: ", i+1)
			fmt.Scan(&grades[i])
		}

		average := calculateAverageGrade(grades)
		fmt.Printf("El promedio de las notas es: %.2f\n", average)
	case 33:
		fmt.Println("[1] Calcular energía cinética")
		fmt.Println("[2] Calcular energía potencial")
		fmt.Println("[3] Calcular energía mecánica")
		fmt.Print("Opción >> ")
		fmt.Scan(&opcion)

		calcularEnergia(opcion)
	case 34:
		var valorInicial, tasaCrecimiento float64
		var periodos int

		fmt.Print("Ingrese el valor inicial >> ")
		fmt.Scan(&valorInicial)
		fmt.Print("Ingrese la tasa de crecimiento (como valor decimal) >> ")
		fmt.Scan(&tasaCrecimiento)
		fmt.Print("Ingrese el número de periodos >> ")
		fmt.Scan(&periodos)

		valorFinal := crecimientoExponencial(valorInicial, tasaCrecimiento, periodos)
		fmt.Printf("El valor final después de %d periodos es %.2f\n", periodos, valorFinal)
	case 35:
		var cantidadPrestamo, interesAnual float64
		var años int

		fmt.Print("Ingrese el monto del prestamo >> ")
		fmt.Scan(&cantidadPrestamo)

		fmt.Print("Ingrese la tasa de interés anual (como decimal) >> ")
		fmt.Scan(&interesAnual)

		fmt.Print("Ingrese el plazo del prestamo en años >> ")
		fmt.Scan(&años)

		pagoMensual := calcularPagoMensual(cantidadPrestamo, interesAnual, años)
		fmt.Printf("El pago mensual es: %.2f\n", pagoMensual)
	default:
		os.Exit(1)
	}
}
