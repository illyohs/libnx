include ./gen/sources.mk

LIBNX_MAJOR := 2
LIBNX_MINOR := 5
LIBNX_PATCH := 0

VERSION := $(LIBNX_MAJOR).$(LIBNX_MINOR).$(LIBNX_PATCH)

TARGET := nx

INCLUDES := ../nx/include ../external/bsd/include
ARCH := -target armv8-a+crc+crypto -mtune=cortex-a57 -fPIC -ftls-model=local-exec

CFLAGS := -g -Wall -Werror -ffunction-sections -fdata-sections $(ARCH) -D__SWITCH__ -DLIBNX_NO_DEPRECATION
CXXFLAGS := $(CFLAGS) -fno-rtti -fno-exceptions -std=gnu++11

export CC=clang
export CXX=clang++
export LD=lld

compile:
	$(CC) $(CFLAGS) -shared $(SRC)