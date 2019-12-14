Определение максимального количества файлов во входной директории:

NodeJS:
Поскольку stream использует основное количество оперативной памяти на считывание в буфер, то количество файлов во входной директории прямо пропорционально размеру буфера (128 в приведенном исполнении лабораторной работы) и количеству оперативной памяти рабочей машины.

Go:
Поскольку goroutine, в отличии от stream в NodeJS, использует дополнительно 4-4.5 КВ оперативной памяти, то максимальное количество файлов во входной директории менее зависимо от размера буфера, в приведенном исполнении лабораторной работы, и прямо пропорционально количеству оперативной памяти рабочей машины (>2 миллиона для 8 GB оперативной памяти, >3 миллиона для 12 GB оперативной памяти).

Дополнительно, стоит отметить, что файловые системы так же имеют свои ограничения по максимальному количеству файлов в одной директории:
FAT32: 65,534 файлов;
NTFS: 4,294,967,295 файлов.