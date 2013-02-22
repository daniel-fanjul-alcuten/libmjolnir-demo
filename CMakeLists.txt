cmake_minimum_required (VERSION 2.8)
project (ice)

add_subdirectory("libfoobar")

include_directories("libfoobar")
set(bazs baz/baz1.c baz/baz2.c baz/baz3.c baz/baz4.c baz/baz5.c baz/baz6.c baz/baz7.c baz/baz8.c baz/baz9.c baz/baz10.c baz/baz11.c baz/baz12.c baz/baz13.c baz/baz14.c baz/baz15.c baz/baz16.c baz/baz17.c baz/baz18.c baz/baz19.c baz/baz20.c baz/baz21.c baz/baz22.c baz/baz23.c baz/baz24.c baz/baz25.c baz/baz26.c baz/baz27.c baz/baz28.c baz/baz29.c baz/baz30.c baz/baz31.c baz/baz32.c baz/baz33.c baz/baz34.c baz/baz35.c baz/baz36.c baz/baz37.c baz/baz38.c baz/baz39.c baz/baz40.c baz/baz41.c baz/baz42.c baz/baz43.c baz/baz44.c baz/baz45.c baz/baz46.c baz/baz47.c baz/baz48.c baz/baz49.c baz/baz50.c)

add_executable(main main.c ${bazs})
target_link_libraries(main foobar)
