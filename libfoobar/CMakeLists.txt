cmake_minimum_required (VERSION 2.8)
project (ice)

include_directories(${CMAKE_CURRENT_SOURCE_DIR})

set(foos foo/foo1.c foo/foo2.c foo/foo3.c foo/foo4.c foo/foo5.c foo/foo6.c foo/foo7.c foo/foo8.c foo/foo9.c foo/foo10.c foo/foo11.c foo/foo12.c foo/foo13.c foo/foo14.c foo/foo15.c foo/foo16.c foo/foo17.c foo/foo18.c foo/foo19.c foo/foo20.c foo/foo21.c foo/foo22.c foo/foo23.c foo/foo24.c foo/foo25.c foo/foo26.c foo/foo27.c foo/foo28.c foo/foo29.c foo/foo30.c foo/foo31.c foo/foo32.c foo/foo33.c foo/foo34.c foo/foo35.c foo/foo36.c foo/foo37.c foo/foo38.c foo/foo39.c foo/foo40.c foo/foo41.c foo/foo42.c foo/foo43.c foo/foo44.c foo/foo45.c foo/foo46.c foo/foo47.c foo/foo48.c foo/foo49.c foo/foo50.c)

set(bars bar/bar1.c bar/bar2.c bar/bar3.c bar/bar4.c bar/bar5.c bar/bar6.c bar/bar7.c bar/bar8.c bar/bar9.c bar/bar10.c bar/bar11.c bar/bar12.c bar/bar13.c bar/bar14.c bar/bar15.c bar/bar16.c bar/bar17.c bar/bar18.c bar/bar19.c bar/bar20.c bar/bar21.c bar/bar22.c bar/bar23.c bar/bar24.c bar/bar25.c bar/bar26.c bar/bar27.c bar/bar28.c bar/bar29.c bar/bar30.c bar/bar31.c bar/bar32.c bar/bar33.c bar/bar34.c bar/bar35.c bar/bar36.c bar/bar37.c bar/bar38.c bar/bar39.c bar/bar40.c bar/bar41.c bar/bar42.c bar/bar43.c bar/bar44.c bar/bar45.c bar/bar46.c bar/bar47.c bar/bar48.c bar/bar49.c bar/bar50.c)

add_library(foobar ${foos} ${bars})
