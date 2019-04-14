package main

import "fmt"

func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}

	quickSort(A, 0, len(A)-1)

	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}

	return 0
}

func quickSort(A []int, start, end int) {
	if start >= end {
		return
	}

	q := partion(A, start, end)
	quickSort(A, start, q-1)
	quickSort(A, q+1, end)
}

func partion(A []int, start, end int) int {
	pivot := A[end]
	q := start
	for i := start; i < len(A)-1; i++ {
		if A[i] < pivot {
			A[i], A[q] = A[q], A[i]
			q++
		}
	}

	A[end], A[q] = A[q], A[end]

	return q
}

func main() {
	fmt.Println(A)
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}))
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
	fmt.Println(largestPerimeter([]int{757, 341, 254, 499, 96, 622, 685, 514, 482, 655, 680, 593, 407, 660, 330, 469, 841, 335, 53, 764, 341, 918, 290, 667, 47, 725, 319, 135, 118, 636, 145, 540, 363, 708, 658, 735, 999, 306, 531, 427, 393, 872, 30, 838, 873, 710, 438, 364, 95, 567, 737, 510, 961, 881, 37, 567, 464, 991, 587, 363, 25, 532, 565, 116, 876, 479, 805, 908, 71, 581, 262, 881, 657, 734, 60, 111, 718, 739, 950, 349, 588, 963, 947, 213, 356, 429, 446, 420, 706, 221, 709, 948, 855, 811, 763, 978, 704, 384, 601, 920, 170, 864, 785, 308, 983, 524, 711, 233, 537, 657, 682, 728, 37, 131, 521, 783, 130, 669, 574, 503, 453, 712, 216, 976, 764, 803, 921, 747, 261, 329, 500, 361, 485, 757, 173, 47, 211, 730, 171, 254, 115, 936, 111, 837, 35, 64, 244, 481, 385, 300, 450, 682, 339, 423, 927, 835, 332, 827, 624, 212, 180, 802, 873, 535, 528, 947, 882, 549, 574, 81, 415, 645, 307, 826, 924, 405, 800, 791, 620, 337, 517, 488, 981, 920, 131, 367, 39, 45, 686, 256, 781, 298, 303, 146, 478, 297, 33, 767, 247, 195, 664, 763, 65, 127, 790, 424, 678, 238, 810, 944, 285, 700, 595, 432, 968, 497, 288, 15, 995, 855, 541, 712, 635, 938, 855, 784, 638, 50, 584, 857, 620, 495, 337, 903, 1000, 407, 687, 520, 406, 447, 586, 209, 53, 869, 383, 80, 311, 928, 545, 916, 16, 219, 517, 226, 31, 201, 830, 753, 324, 989, 701, 976, 713, 228, 227, 391, 268, 301, 674, 690, 994, 224, 851, 638, 117, 341, 146, 235, 676, 172, 429, 744, 806, 899, 549, 633, 148, 181, 948, 381, 861, 93, 143, 974, 366, 901, 517, 161, 166, 492, 355, 441, 388, 407, 17, 34, 472, 592, 32, 184, 388, 156, 400, 764, 435, 968, 798, 130, 277, 687, 955, 476, 135, 644, 313, 303, 28, 635, 667, 817, 651, 341, 98, 637, 870, 141, 695, 425, 997, 467, 961, 622, 624, 987, 913, 256, 781, 338, 304, 144, 253, 137, 860, 377, 162, 476, 632, 181, 616, 671, 486, 900, 447, 697, 389, 787, 858, 495, 574, 329, 583, 423, 784, 382, 715, 939, 515, 476, 710, 558, 872, 52, 133, 524, 240, 792, 506, 958, 738, 138, 378, 836, 469, 238, 659, 54, 346, 458, 971, 536, 44, 34, 494, 224, 573, 443, 90, 804, 415, 90, 48, 690, 713, 302, 243, 276, 636, 611, 226, 933, 788, 147, 94, 691, 453, 414, 518, 322, 939, 961, 551, 338, 79, 1000, 186, 671, 404, 675, 653, 965, 218, 268, 410, 656, 793, 615, 884, 911, 55, 360, 297, 526, 913, 120, 301, 551, 604, 454, 545, 531, 3, 642, 85, 67, 394, 235, 731, 862, 599, 804, 574, 951, 981, 855, 226, 19, 336, 259, 441, 342, 305, 250, 727, 78, 985, 672, 110, 78, 719, 57, 375, 449, 750, 593, 497, 376, 930, 67, 822, 750, 10, 170, 371, 465, 5, 542, 965, 388, 508, 888, 873, 337, 590, 863, 425, 812, 347, 41, 306, 813, 413, 343, 613, 12, 978, 81, 227, 78, 560, 164, 921, 847, 279, 113, 663, 290, 474, 395, 390, 241, 206, 804, 568, 831, 391, 83, 501, 222, 823, 171, 1000, 290, 513, 343, 706, 444, 303, 606, 217, 25, 290, 619, 702, 935, 699, 27, 774, 916, 531, 682, 896, 834, 14, 729, 454, 372, 542, 189, 960, 958, 508, 710, 199, 981, 207, 158, 215, 904, 322, 302, 436, 922, 449, 627, 370, 409, 33, 138, 85, 470, 73, 177, 724, 723, 441, 278, 910, 2164, 14, 217, 310, 65, 606, 488, 477, 120, 894, 495, 149, 891, 481, 716, 297, 912, 89, 981, 414, 715, 398, 844, 890, 746, 970, 767, 821, 760, 291, 876, 423, 869, 845, 580, 959, 265, 882, 217, 966, 537, 526, 557, 911, 662, 991, 579, 358, 636, 405, 537, 913, 441, 267, 892, 376, 657, 627, 99, 664, 359, 549, 941, 520, 65, 759, 550, 198, 207, 284, 887, 853, 581, 23, 364, 394, 754, 936, 532, 480, 16203, 965, 531, 429, 72, 963, 640, 365, 797, 349, 219, 235, 571, 282, 732, 976, 625, 277, 298, 726, 619, 668, 917, 598, 895, 489, 128, 365, 701, 537, 68, 27, 32, 546, 169, 323, 548, 38, 819, 498, 276, 498, 520, 959, 95, 202, 615, 387, 451, 429, 255, 572, 19, 61, 496, 584, 142, 977, 1531, 45, 874, 442, 288, 66, 938, 929, 532, 786, 66, 734, 607, 405, 502, 597, 241, 669, 758, 445, 955, 708, 894, 875, 798, 993, 818, 336, 215, 631, 574, 661, 422, 489, 60, 23, 90, 51, 41, 526, 489, 231, 860, 635, 345, 865, 519, 350, 426, 755, 457, 999, 841, 942, 336, 943, 414, 526, 262, 393, 749, 967, 55, 115, 619, 446, 44, 184, 804, 910, 737, 599, 891, 851, 752, 327, 1000, 977, 543, 942, 985, 757, 106, 411, 343, 370, 623, 34, 625, 60, 754, 601, 434, 94, 187, 716, 451, 310, 379, 700, 362, 543, 528, 389, 454, 950, 83, 508, 729, 503, 688, 1000, 787, 3891, 903, 105, 482, 750, 843, 969, 841, 478, 136, 971, 260, 439, 753, 232, 233, 398, 134, 885, 342, 209, 253, 426, 656, 154, 299, 567, 74, 136, 630, 979, 884, 725, 472, 943, 78, 289, 962, 173, 535, 318, 157, 939, 39, 976, 443, 15, 107, 259, 269, 146, 690, 79, 506, 523, 546, 170, 776, 736, 721, 569, 962, 553, 507, 6, 311, 527, 300, 997, 553, 174, 877, 244, 540, 818, 673, 980, 518, 913, 257, 82, 962, 264, 130, 79, 425, 871, 808, 823, 627, 928, 461, 709, 218, 740, 822, 665, 729, 994, 614, 298, 36, 82, 601, 409, 488, 384, 869, 818, 74, 255, 11, 512, 75, 572, 78, 320, 868, 803, 49, 277, 439, 177, 378, 647, 290, 280, 923, 776, 622, 203, 523, 824, 179, 68, 867, 328, 646, 145, 114, 496, 234, 599, 335, 645, 221, 557, 561, 797, 188, 688, 697, 351, 91, 965, 285, 485, 207, 774, 33, 87, 167, 926, 249, 906, 989, 523, 985, 14, 188, 64, 507, 205, 450, 725, 282, 891, 677, 84, 955, 754, 810, 91, 492, 910, 383, 863, 941, 812, 859, 514, 757, 54, 126, 45, 413, 938, 647, 809, 620, 307, 225, 628, 41, 113, 757, 552, 578, 877, 607, 955, 192, 937, 901, 146, 815, 67, 144, 806, 240, 199, 705, 650, 60, 976, 917, 694, 548, 962, 985, 255, 785, 939, 530, 997, 400, 4, 832, 353, 202, 734, 605, 408, 927, 387, 947, 34, 712, 516, 896, 508, 513, 81, 862, 960, 931, 111667, 838, 569, 866, 270, 290, 396, 24, 496, 426, 826, 358, 859, 475, 88, 510, 902, 748, 648, 58, 351, 710, 885, 173, 247, 937, 874, 597, 487, 53, 365, 498, 352, 869, 659, 458, 50, 774, 251, 672, 930, 864, 367, 675, 548, 922, 860, 769, 472, 73, 933, 29, 259, 412, 161, 380, 82, 153, 832, 948, 257, 521, 671, 883, 173, 388, 639, 313, 485, 54, 951, 818, 994, 81, 51, 905, 959, 231, 543, 730, 954, 672, 73, 562, 157, 263, 10003, 596, 834, 690, 210, 231, 185, 180, 581, 883, 459, 238, 208, 98, 988, 418, 806, 321, 985, 183, 749, 427, 538, 491, 51, 74, 306, 722, 28, 399, 268, 611, 441, 10, 321, 849, 733, 228, 735, 411, 883, 345, 751, 753, 896, 91, 627, 737, 969, 529, 274, 56, 168, 716, 789, 132, 176, 616, 69, 319, 525, 284, 530, 68, 23, 214, 617, 628, 68, 718, 526, 408, 851, 636, 743, 142, 237, 930, 110, 199, 158, 598, 323, 847, 621, 392, 982, 112, 515, 309, 504, 129, 915, 328, 328, 677, 722, 427, 14, 608, 737, 865, 469, 961, 465, 413, 803, 875, 161, 561, 497, 114, 786, 293, 436, 999, 247, 410, 845, 575, 444, 454, 692, 746, 965, 423, 711, 579, 614, 405, 643, 132, 466, 347, 152, 474, 215, 645, 369, 369, 576, 603, 305, 807, 799, 720, 958, 845, 904, 292, 91, 764, 82, 99, 408, 816, 260, 982, 595, 437, 301, 261, 438, 871, 429, 864, 538, 765, 702, 380, 792, 713, 692, 103, 668, 554, 106, 448, 41, 514, 668, 847, 68, 812, 414, 666, 825, 548, 882, 531, 65, 260, 317, 85, 180772, 266, 585, 210, 13, 437, 436, 792, 653, 494, 156, 845, 29, 493, 994, 225, 184, 397, 48, 733, 851, 496, 42571, 34, 103, 686, 638, 666, 604, 99, 615, 481, 752, 508, 129, 445, 486, 528, 412, 803, 128, 63, 317, 796, 750, 394, 654, 739, 528, 869, 831, 341, 713, 638, 229, 329, 548, 641, 907, 288, 326, 202, 875, 867, 506, 494, 637, 840, 486, 150, 996, 991, 293, 170, 842, 790, 471, 707, 298, 269, 768, 961, 219, 679, 968, 955, 908, 458, 59, 640, 704, 785, 524, 427, 679, 857, 836, 917, 154, 24, 485, 565, 30, 776, 688, 54, 161, 158, 433, 211, 220, 947, 414, 819, 611, 873, 315, 334, 690, 888, 63, 487, 135, 608, 589, 121, 761, 394, 369, 177, 197, 101, 192, 363, 62, 330, 198, 485, 930, 693, 161, 214, 91, 395, 313, 1000, 949, 592, 65, 28, 886, 393, 359, 188, 698, 532, 824, 753, 506, 190, 409, 363, 13, 932, 630, 734, 84, 789, 523, 512, 593, 814, 324, 705, 814, 104, 579, 162, 648, 673, 974, 507, 882, 700, 714, 892, 41, 103, 129, 912, 491, 891, 5, 413, 233, 442, 997, 469, 141, 695, 680, 357, 449, 874, 421, 2, 351, 570, 319, 69, 814, 266, 540, 165, 443, 618, 980, 854, 18, 941, 565, 943, 110, 977, 255, 889, 810, 239, 216, 430, 566, 760, 236, 581, 808, 483, 926, 68, 368, 990, 987, 604, 926, 487, 197, 935, 502, 405, 639, 727, 385, 906, 762, 460, 148, 727, 270, 175, 546, 735, 160, 459, 828, 735, 915, 326, 102, 772, 405, 208, 35, 496, 691, 187, 851, 543, 63, 13, 393, 857, 14, 399, 668, 942, 554, 910, 411, 952, 620, 760, 103, 669, 416, 985, 408, 787, 830, 666, 570, 950, 944, 723, 791, 422, 448, 565, 549, 819, 221, 282, 872, 47, 850, 667, 159, 102, 223, 831, 330, 24, 470, 701, 312, 787, 111, 281, 432, 967, 551, 831, 604, 875, 467, 533, 773, 196, 365, 363, 688, 885, 951, 391, 370, 648, 824, 600, 177, 477, 811, 290, 882, 297, 128, 862, 714, 99, 578, 905, 214, 920, 810, 426, 915, 786, 233, 743, 115, 436, 559, 137, 651, 172, 691, 845, 257, 733, 179, 991, 730, 345, 691, 534, 220, 128, 143, 265, 156, 711, 833, 193, 301, 129, 407, 819, 480, 42, 992, 474, 746, 46, 230, 38, 208, 579, 886, 854, 62, 782, 469, 201, 709, 334, 214, 400, 294, 211, 933, 848, 503, 468, 179, 702, 452, 645, 569, 82, 965, 833, 33, 962, 941, 527, 467, 536, 243, 576, 775, 353, 954, 940, 984, 667, 104, 160, 467, 607, 703, 451, 68, 333, 780, 29, 912, 373, 993, 695, 845, 682, 176, 29, 869, 42, 846, 898, 144, 60, 327, 283, 568, 746, 498, 176, 251, 933, 37, 832, 189, 475, 983, 156, 836, 522, 579, 637, 249, 294, 609, 115, 634, 492, 764, 381, 831, 698, 682, 410, 854, 509, 754, 116, 857, 312, 774, 959, 918, 943, 8, 573, 733, 241, 831, 319, 783, 193, 510, 943, 455, 977, 899, 109, 741, 37, 90, 47, 258, 219, 150, 750, 91, 473, 440, 768, 21, 864, 125, 404, 888, 120, 552, 551, 576, 704, 357, 896, 516, 156, 952, 577, 900, 168, 931, 731, 893, 628, 524, 66, 884, 292525, 547, 735, 899, 975, 498, 406, 151, 564, 827, 363, 862, 607, 130, 573, 793, 332, 818, 762, 418, 988, 288, 480, 141, 319, 678, 278, 516, 105, 849, 319, 961, 908, 751, 118, 825, 132, 757, 631, 294, 156, 586, 704, 42, 837, 464, 807, 52, 454, 818, 58, 491, 654, 113, 218, 959, 488, 868, 799, 131, 154, 192, 390, 199, 424, 46, 782, 540, 108, 558, 403, 327, 312, 726, 184, 805, 194, 423, 383, 260, 242, 9, 723, 142, 608, 579, 879, 560, 378, 419, 452, 992, 626, 496, 739, 148, 265, 791, 966, 937, 623, 754, 166, 30, 338, 35, 48, 322, 868, 552, 820, 354, 70, 471, 350, 746, 469, 53, 705, 870, 22, 72, 227, 856, 311, 33, 282, 697, 400, 25, 472, 967, 207, 863, 753, 583, 569, 954, 608, 768, 579, 462, 473, 154, 581, 234, 83, 467, 386, 64, 731, 446, 91, 610, 856, 880, 770, 446, 3, 199, 413, 384, 476, 427, 574, 641, 455, 46, 221, 565, 985, 457, 971, 805, 248, 899, 195, 50, 885, 872, 27, 112, 545, 696, 537, 388, 698, 612, 647, 923, 399, 177, 920, 621, 627, 768, 442, 684, 288, 414, 537, 573, 231, 476, 356, 222, 923, 759, 633, 858, 247, 452, 273, 632, 133, 449, 504, 859, 826, 799, 728, 252, 665, 801, 453, 423, 247, 401, 607, 918, 967, 460, 328, 318, 141, 34, 572, 899, 392, 712, 584, 919, 294, 924, 229, 73, 675, 1, 193, 257, 908, 388, 448, 45, 726, 510, 652, 608, 552, 258, 723, 145, 608, 603, 644, 925, 1, 367, 62, 712, 740, 234, 37, 370, 875, 500, 969, 107, 657, 358, 201, 595, 937, 155, 972, 584, 462, 61, 798, 917, 500, 993, 346, 60, 29, 854, 504, 121, 949, 218, 301, 196, 819, 747, 66, 285, 215, 89, 233, 177, 123, 717, 100, 568, 376, 842, 563, 996, 878, 5, 886, 198, 717, 659, 532, 695, 150, 439, 789, 433, 871, 80, 680, 139, 557, 8, 561, 148, 678, 90, 480, 321, 93, 130, 666, 25, 6102, 428, 945, 947, 682, 128, 499, 172, 423, 511, 72, 370, 548, 599, 66, 879, 678, 559, 784, 4, 22, 745, 660, 812, 599, 218, 683, 638, 411, 227, 405, 387, 323, 568, 785, 637, 626, 211, 193, 142, 517, 525, 812, 159, 725, 214, 165, 505, 173, 489, 249, 789, 394, 944, 934, 451, 697, 660, 100, 259, 486, 626, 872, 183, 102, 24, 152, 27, 63, 676, 928, 17, 699, 859, 227, 500, 988, 106, 723, 284, 792, 685, 829, 847, 912, 547, 306, 195, 584, 256, 413, 40, 656, 903, 53, 182, 751, 403, 225, 715, 218, 195, 726, 79, 222, 497, 519, 769, 64, 651, 413, 570, 682, 232, 144, 619, 688, 51, 739, 694, 402, 719, 393, 639, 661, 926, 462, 654, 274, 163, 165, 107, 631, 589, 758, 944, 647, 402, 201, 826, 193, 907, 966, 961, 155, 829, 645, 865, 543, 68931, 643, 693, 484, 894, 823, 423, 764, 644, 70, 419, 465, 383, 755, 701, 517, 610, 323, 883, 808, 183, 562, 712, 401, 918, 848, 602, 648, 740, 115, 542, 947, 362, 712, 115, 565, 803, 821, 499, 356, 674, 436, 396, 831, 933, 444, 766078, 380, 782, 78, 809, 781, 514, 811, 976, 530, 213, 705, 951, 648, 882, 229, 573, 645, 313, 495, 676, 473, 375, 861, 66, 312, 213, 877, 366, 524, 763, 260, 441, 768, 694, 854, 328, 127, 531, 993, 293, 172, 312, 554, 414, 688, 326, 62, 319, 3, 18, 920, 633, 422, 153, 626, 563, 6, 483, 880, 414, 999, 278, 245, 727, 613, 462, 948, 396, 204, 452, 24, 265, 162, 241, 335, 960, 413, 947, 235, 509, 26, 90, 339, 805, 471, 630, 592, 802, 897, 100, 605, 664, 755, 196, 956, 408, 262, 918, 380, 803, 332, 386, 786, 51, 562, 743, 42, 802, 467, 133, 908, 492, 610, 861, 308, 846, 838, 980, 577, 560, 473489, 734, 827, 166, 738, 998, 329, 364, 783, 557, 686, 81, 41, 933, 652, 378, 648, 268, 294, 378, 331, 257, 822, 807, 460, 153, 207, 785, 868, 704, 295, 296, 841, 104, 391, 382, 452, 322, 430, 235, 239, 2, 297, 282, 825, 386, 236, 918, 323, 920, 51, 993, 330, 829, 28, 560, 683, 784, 916, 113, 634, 143, 864, 758, 598, 77, 965, 511, 767, 121, 202, 908, 887, 436, 483, 992, 82, 431, 775, 352, 551, 616, 738, 195, 821, 859, 17, 931, 616, 509, 557, 978, 261, 663, 783, 188, 567, 872, 983, 434, 924, 957, 474, 327, 642, 422, 416, 71, 651, 807, 982, 679, 433, 57, 408, 233, 423, 969, 963, 595, 484, 999, 787, 656, 609, 173, 568, 432, 874, 314, 624, 691, 991, 385, 397, 730, 127, 809, 152, 291, 93, 753, 620, 341, 503, 349, 265, 698, 600, 832, 577, 320, 830, 778, 900, 692, 525, 517, 230, 790, 958, 293, 534, 923, 147, 57, 529, 81, 471, 837, 934, 957, 34, 189, 117, 592, 400, 907, 482, 76, 810, 804, 411, 255, 132, 974, 833, 334, 484, 393, 101, 657, 744, 829, 345, 844, 313, 640, 725, 944, 111, 526, 737, 777, 574, 276, 765, 954, 210, 630, 678, 95, 54, 510, 329, 849, 532, 910, 404, 516, 825, 591, 204, 578, 384, 395, 966, 923, 618, 674, 308, 896, 920, 853, 492, 268, 934, 847, 615, 787, 193, 800, 67, 127, 180, 582, 403, 934, 762, 591, 699, 657, 772, 823, 127, 81, 666, 117, 328, 460, 902, 549, 489, 704, 774, 704, 185, 111, 517, 640, 872, 643, 868, 896, 483, 735, 561, 619, 174, 408, 645, 962, 863, 304, 68, 995, 336, 856, 210, 439, 373, 935, 340, 861, 755, 182, 543, 743, 342, 364, 975, 964, 82, 81, 880, 865, 381, 488, 653, 522, 539, 542, 160, 562, 933, 759, 96, 43, 989, 337, 627, 409, 594, 446, 187, 158, 956, 658, 628, 390, 308, 472, 178, 640, 283, 226, 364, 187, 16, 516, 590, 434, 260, 541, 733, 621, 294, 69, 664, 609, 565, 842, 348, 368, 961, 532, 544, 957, 974, 158, 455, 90, 328, 759, 735, 949, 881, 290, 523, 4, 804, 103, 668, 26317, 971, 772, 450, 362, 801, 317, 358, 946, 191, 674, 212, 72, 235}))
}
