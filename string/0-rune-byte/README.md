# Объяснение / Решение

Отрывок из документации
```go
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32
```

Байты - это то же самое, что и uint8. Они занимают в памяти 8 битов, 
и могут хранить числовые значения до 256.

Руны - это то же самое, что и int32. Они занимают в памяти 32 битов,
и могут хранить числовые значения до

Строки(string) в golang - это массив байтов, именно байтов. Поэтому,
длина строки будет равна 10, так как 10 байтов.

1 руна - это 4 байта