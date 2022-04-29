.code32 

.global _start
_start:

#zero 4 pages
 xor %eax, %eax
 mov $0x1000,%edi
 mov $0x5000, %ecx
 rep stosb

#P4ML[0] -> 0x2000 (PDPT-A)
  mov $$(0x3000|3),%eax
  mov %eax,0x1FF4

#P4ML[1] -> 0x2000 (PDPT-B)
mov $(0x3000|3),%eax
mov %eax,0x2000

# PDPT-A[0] -> 0x4000(PD)
mov $(0x3000|3),%eax
mov %eax,0x4FF0


# PD[0..511] -> 0.1022MB
  mov $0x90,%eax
  mov $0x40000,%ebx
  mov $512,%ecx