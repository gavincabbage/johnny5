# johnny5

include Makefile.inc

CORE_DIR=core
WEB_DIR=web
FEED_DIR=feed

.PHONY: core web feed package clean

core:
	$(MAKE) -C $(CORE_DIR)

web:
	$(MAKE) -C $(WEB_DIR)

feed:
	$(MAKE) -C $(FEED_DIR)

pack: core web #feed
	$(TARZIP) $(TARGET) $(TARGET_DIR)

all: core web pack #feed

clean:
	rm -rf $(TARGET_DIR)
	rm -f $(TARGET)
