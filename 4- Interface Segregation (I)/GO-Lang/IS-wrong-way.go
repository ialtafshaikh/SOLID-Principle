type I1 interface { // consumed by C1
    M1()
    M2()
    M3()       
}
type I2 interface { // consumed by C2 and C3
    M3()       
    M4()
}