# Lulu
ð§  Lulu(é²é²) æ¯ä¸ä¸ªæä»¶åç¼©å¤ççå·¥å·ç®±ã
> ä½¿ç¨åï¼å¡å¿åå¨æµè¯æå¡å¨å°è¯å®è£ libvips ã

## ä¾èµ
- [libvips](https://www.libvips.org/install.html) 
   - â­ï¸ è¿éæä¾å¦ä¸ç§ libvips çå®è£æ¹å¼ï¼éè¿æå¨æå®ä¾èµï¼è¿æ ·å¯ä»¥é¿å libvips çå®è£åºç°ä¸å¯æ§çé®é¢ï¼è¦çé¨å C åºç¡ç»ä»¶ï¼ã
      1. å°æ¬ä»åºç libvips.so.42.13.3 å­æ¾è³ä»»æä½ç½®ï¼åè®¾å­æ¾å¨ /tmp
      3. cd /etc
      4. vi ld.so.confï¼æå® so æä»¶çå­æ¾ç®å½ï¼/tmpï¼
      5. cd/sbin
      6. ldconfig -v

## æ¯æ
- å¾çåç¼©

original before size 1.1 MB (1,070,107 bytes)  | compressed after 61 KB (62,060 bytes)
------------- | -------------
![screenshot](https://s3.bmp.ovh/imgs/2021/11/f9fec302580982cc.jpeg "compress example")  | ![screenshot](https://s3.bmp.ovh/imgs/2021/11/8cc7d3099d517116.jpeg "compress example")


```
$lulu image -h

åç¼©æå®æä»¶å¤¹ä¸çå¾çï¼æ¯æå¤çº§æä»¶å¤¹ã ä¾å¦: lulu image

Usage:
  lulu image [flags]

Flags:
  -d, --dir-path string   æå®å­æ¾å¾ççæä»¶å¤¹ï¼æ¯æå¨å¤çº§æä»¶å¤¹ä¸è·åå¾çã
  -q, --quality string    æå®åç¼©è´¨éï¼åç¼©è´¨éèå´ä¸º 1(æä½è´¨é) - 100(æé«è´¨é)ï¼è´¨éè¶ä½ï¼å¾çå¤§å°è¶å°ã
  -s, --suffix string     æå®å¾å¤çå¾ççåç¼ï¼ç®åæ¯æ pngï¼jpegï¼jpgã
```

![ææ](https://s3.bmp.ovh/imgs/2021/11/de561b2905c114c4.jpg)

## ç¼è¯
äº¤åç¼è¯ï¼å¯¹æèè¨æ¯éå¸¸ææçãæ¨èå¨ç¯å¢ä¸­ï¼å®è£ libvipsã

## å¸¸è§é®é¢

- libvips

  ```
  # pkg-config --cflags vips vips vips vips
  Package vips was not found in the pkg-config search path.
  Perhaps you should add the directory containing `vips.pc'
  to the PKG_CONFIG_PATH environment variable
  No package 'vips' found
  ```
  è§£å³æ¹æ¡ï¼https://www.cxybb.com/article/a_c12345/101725160
