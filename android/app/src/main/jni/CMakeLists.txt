cmake_minimum_required(VERSION 3.13)

# Define the library name here.
project(appmodules)

# This file includes all the necessary to let you build your React Native application
include(${REACT_ANDROID_DIR}/cmake-utils/ReactNative-application.cmake)

# Define where the additional source code lives. We need to crawl back the jni, main, src, app, android folders
target_sources(${CMAKE_PROJECT_NAME} PRIVATE ../../../../../shared/NativeChatModule.cpp)

# Define where CMake can find the additional header files. We need to crawl back the jni, main, src, app, android folders
target_include_directories(${CMAKE_PROJECT_NAME} PUBLIC ../../../../../shared)

if(NOT DEFINED GO_LIB_DIR)
  message(FATAL_ERROR
        "GO_LIB_DIR not set Build through Gradle so android/app/build.gradle "
        "can pass -DGO_LIB_DIR=<app>/build/go.")
endif()

set(GO_ABI_DIR "${GO_LIB_DIR}/${ANDROID_ABI}")

add_library(chat_go SHARED IMPORTED)
set_target_properties(chat_go PROPERTIES
    IMPORTED_LOCATION "${GO_ABI_DIR}/libchat.so"
    IMPORTED_SONAME "libchat.so")

target_include_directories(${CMAKE_PROJECT_NAME} PUBLIC "${GO_ABI_DIR}")
target_link_libraries(${CMAKE_PROJECT_NAME} chat_go)
