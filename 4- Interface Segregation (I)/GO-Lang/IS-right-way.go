type I1 interface { // consumed by C1
    M1()
    M2()
    M3()       // also defined in I2
}
type I2 interface { // consumed by C2 and C3
    M3()       // here M3 not in a separate interface as no client
               // use an interface with only M3()
    M4()
}
type I3 interface { // consumed by C4
    M5()       // similarly M5() only used along with I1 and I2
               // thus not needed to have it in a separate interface
    I1
    I2
}