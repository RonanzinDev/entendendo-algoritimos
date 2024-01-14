module main

fn regresiva(i int) {
	println(i)
	if i <= 1 {
		return
	} else {
		regresiva(i - 1)
	}
}

fn main() {
	regresiva(10)
}
