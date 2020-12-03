# Tarea2SD_20201

Integrantes:
-   Andrés Herera
    ROL: 201473593-6
-   Sebastián Muñoz
    ROL: 201473503-0

Comandos MAKEFILE:
make runNamenode        ejecuta el servidor del namenode, debe ejecutarse solo en la maquina dist144
make runDatanode        ejecuta el servidor del datanode, ejecutarse en maquinas dist141, dist142 y dist143
make runCliente         ejecuta al cliente, puede ejecutarse en cualquier maquina

Si el datanode de la maquina dist141 no esta levantada, el programa no funciona.

Programa SOLO recibe archivos sin espacios.
Cliente elige algoritmo de exclusion a utilizar.
Archivos de prueba por el lado del cliente se dejan en la carpeta raiz.
Fragmentos de los archivos quedan en la carpeta libros
Cada vez que se quiera realizar una funcion de cliente, se debe levantar el programa del cliente.

Casos nodo no disponible:
    -Para todas las maquinas, existe un 10% de probabilidad de no estar disponible
    -Nodo no alcanzable

La carpeta "libros" debe existir la hora de ejecutarse el proceso.

La distribucion para ejecutar la tarea es la siguiente:
dist141: datanode1
dist142: datanode2
dist143: datanode3
dist144: namenode