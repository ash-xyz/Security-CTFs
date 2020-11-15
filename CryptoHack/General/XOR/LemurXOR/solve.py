
import cv2

img1 = cv2.imread('flag.png')
img2 = cv2.imread('lemur.png')

flag = cv2.bitwise_xor(img2, img1, mask=None)
cv2.imwrite('dec_flag.png', flag)
