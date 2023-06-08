package utils

type TestCase []struct {
	input  string
	output string
}

var sample = TestCase{
	{"1 2", "3"},
	{"-5 8", "3"},
	{"10 -7", "3"},
	{"10 -7", "3"},
	{"10 -7", "3"},
	{"10 -7", "3"},
	{"10 -7", "3"},
	{"10 -7", "3"},
	{"10 -7", "3"},
	{"9999999 9999999", "19999998"},
}
var T1 = TestCase{
	// {"3\nJohn 25\nSarah 20\nTom 30\n", "(Tom, 30)(John, 25)(Sarah, 20)"},
	// {"1\nAdam 18\n", "(Adam, 18)"},
	// {"2\nAmy 22\nBob 22\n", "(Amy, 22)(Bob, 22)"},
	// {"4\nSara 40\nDavid 25\nPeter 30\nMary 35\n", "(Sara, 40)(Mary, 35)(Peter, 30)(David, 25)"},
	// {"0\n", "()"},
	// {"5\nAlice 27\nTom 33\nEmma 29\nHenry 24\nOlivia 26\n", "(Tom, 33)(Emma, 29)(Alice, 27)(Olivia, 26)(Henry, 24)"},
	// {"2\nJohn 30\nMary 30\n", "(John, 30)(Mary, 30)"},
	// {"3\nJames 22\nLily 22\nMark 22\n", "(James, 22)(Mark, 22)(Lily, 22)"},
	// {"2\nNathan 19\nEthan 18\n", "(Nathan, 19)(Ethan, 18)"},
	// {"4\nSophia 28\nWilliam 31\nLucas 28\nMia 27\n", "(William, 31)(Lucas, 28)(Sophia, 28)(Mia, 27)"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"7\nSintia 40\nJames 34\nUdin 45\nDani 60\nFarhan 47\nAdam 50\nIbrahim 45", "Berat badan Dani tidak memenuhi syarat\nIbrahim tidak mendapatkan antrian\nSintia 40\nJames 34\nUdin 45\nFarhan 47\nAdam 50"},
	{"3\nLily 49\nLucas 48\nChloe 47", "Lily 49\nLucas 48\nChloe 47"},
	{"3\nLily 49\nLucas 48\nChloe 47", "Lily 49\nLucas 48\nChloe 47"},
	{"0", ""},
}
var T2 = TestCase{
	// {"5\n1 2 3 4 5", "5 4 3 2 1 \n1 2 3 4 5"},
	// {"3\n10 20 30", "30 20 10 \n10 20 30"},
	// {"0", ""},
	// {"1\n100", "100 \n100"},
	// {"7\n9 8 7 6 5 4 3", "3 4 5 6 7 8 9 \n9 8 7 6 5 4 3"},
	// {"6\n-5 -4 -3 -2 -1 0", "0 -1 -2 -3 -4 -5 \n-5 -4 -3 -2 -1 0"},
	// {"4\n0 1 0 1", "1 0 1 0 \n0 1 0 1"},
	// {"2\n3 7", "7 3 \n3 7"},
	// {"8\n5 5 5 5 5 5 5 5", "5 5 5 5 5 5 5 5 \n5 5 5 5 5 5 5 5"},
	// {"6\n-10 -20 -30 0 30 20", "20 30 0 -30 -20 -10 \n-10 -20 -30 0 30 20"},
	{"0", "0"},
	{"1", "1"},
	{"5", "5"},
	{"8", "21"},
	{"10", "55"},
	{"15", "610"},
	{"20", "6765"},
	{"25", "75025"},
	{"29", "514229"},
	{"30", "832040"},
}
var T3 = TestCase{
	// {"97533 71606 4687 87802 6710 68323 68381 56850 55018 83829 10449 77437 9001 20223 209 50665 4710 79754 73812 80141 76659 46663 11645 37595 1775 8336 46328 4718 64906 567 44903 6599 9906 97917 94619 64367 21378 24392 35776 97078 27107 10518 46618 94446 25470 9776 4807 97220 95223 95563 79958 20731 61731 28627 4705 85997 72747 98065 23997 42994 13905 72198 24247 52875 63685 48827 87080 64281 63591 94846 42177 7336 98346 63663 31029 40949 30204 10932 31472 43307 12723 72384 78407 55598 35386 30326 58173 11898 64118 46018 64508 80220 49561 46891 97027 43646 13848 82362 23550 4158 x", "4158 23550 82362 13848 43646 97027 46891 49561 80220 64508 46018 64118 11898 58173 30326 35386 55598 78407 72384 12723 43307 31472 10932 30204 40949 31029 63663 98346 7336 42177 94846 63591 64281 87080 48827 63685 52875 24247 72198 13905 42994 23997 98065 72747 85997 4705 28627 61731 20731 79958 95563 95223 97220 4807 9776 25470 94446 46618 10518 27107 97078 35776 24392 21378 64367 94619 97917 9906 6599 44903 567 64906 4718 46328 8336 1775 37595 11645 46663 76659 80141 73812 79754 4710 50665 209 20223 9001 77437 10449 83829 55018 56850 68381 68323 6710 87802 4687 71606 97533"},
	// {"79115 96396 35797 12002 97398 5694 67056 79422 87957 19935 73138 x", "73138 19935 87957 79422 67056 5694 97398 12002 35797 96396 79115"},
	// {"50065 12257 63964 40815 1063 47315 66817 4786 85134 35198 7562 4632 73679 25628 88874 10711 77481 94071 99363 83782 72122 4407 29460 59802 42072 51849 2470 87953 4956 63666 86215 27488 78545 27126 43719 55830 421 44283 9142 74289 76687 48491 32574 36185 64250 56293 85045 97484 29278 89577 82426 70764 64763 6747 58187 31945 25100 30940 63641 89876 87902 61828 5381 34174 51065 7821 6496 13844 80386 80402 42544 74419 91648 91818 50130 71803 90211 68827 55275 67206 10196 32699 96298 63020 60364 12906 42295 34896 3473 26613 15166 x", "15166 26613 3473 34896 42295 12906 60364 63020 96298 32699 10196 67206 55275 68827 90211 71803 50130 91818 91648 74419 42544 80402 80386 13844 6496 7821 51065 34174 5381 61828 87902 89876 63641 30940 25100 31945 58187 6747 64763 70764 82426 89577 29278 97484 85045 56293 64250 36185 32574 48491 76687 74289 9142 44283 421 55830 43719 27126 78545 27488 86215 63666 4956 87953 2470 51849 42072 59802 29460 4407 72122 83782 99363 94071 77481 10711 88874 25628 73679 4632 7562 35198 85134 4786 66817 47315 1063 40815 63964 12257 50065"},
	// {"39704 23299 48902 27625 27509 58595 74840 6204 57090 94192 99068 3423 28001 33537 48647 44513 69961 74914 6022 95461 15967 44922 76097 76471 86180 65859 15495 32947 82963 25710 17971 14552 4041 63951 43525 20073 16924 30737 11684 87104 80091 92991 33227 32393 61001 21147 1829 20547 91787 90590 69423 x", "69423 90590 91787 20547 1829 21147 61001 32393 33227 92991 80091 87104 11684 30737 16924 20073 43525 63951 4041 14552 17971 25710 82963 32947 15495 65859 86180 76471 76097 44922 15967 95461 6022 74914 69961 44513 48647 33537 28001 3423 99068 94192 57090 6204 74840 58595 27509 27625 48902 23299 39704"},
	// {"52125 59888 89498 65597 32247 63003 95317 34956 2620 92022 39360 86884 5623 93206 64348 7395 71397 48728 2706 85202 72223 76063 66606 65235 43896 88049 73928 22910 27663 27379 62839 88560 69785 22612 66232 44411 75103 x", "75103 44411 66232 22612 69785 88560 62839 27379 27663 22910 73928 88049 43896 65235 66606 76063 72223 85202 2706 48728 71397 7395 64348 93206 5623 86884 39360 92022 2620 34956 95317 63003 32247 65597 89498 59888 52125"},
	// {"95185 70660 47505 56132 99236 1830 47360 17842 84412 82784 86693 55674 12850 70745 2340 30685 71711 46220 61931 40590 59901 39059 74819 88504 94305 6891 54289 88938 84765 28720 1188 75495 94686 62183 83429 80708 61197 87649 76579 14822 58410 93624 98860 94052 16239 63220 66124 14053 52764 45841 7768 70356 49896 99634 64940 47498 91445 57356 19337 96618 15675 42214 68591 46658 72522 44479 16037 50052 76467 41400 60501 1159 62784 53556 71256 82145 2670 85655 97780 44110 11718 33607 33074 52727 63338 x", "63338 52727 33074 33607 11718 44110 97780 85655 2670 82145 71256 53556 62784 1159 60501 41400 76467 50052 16037 44479 72522 46658 68591 42214 15675 96618 19337 57356 91445 47498 64940 99634 49896 70356 7768 45841 52764 14053 66124 63220 16239 94052 98860 93624 58410 14822 76579 87649 61197 80708 83429 62183 94686 75495 1188 28720 84765 88938 54289 6891 94305 88504 74819 39059 59901 40590 61931 46220 71711 30685 2340 70745 12850 55674 86693 82784 84412 17842 47360 1830 99236 56132 47505 70660 95185"},
	// {"78614 34046 4370 36455 53012 2090 40942 86444 72309 98828 67544 98564 74877 94473 77884 68095 31961 25809 27028 65246 54904 297 86751 65445 8003 29835 60988 20377 82456 52447 33419 39091 38489 42742 78506 24621 51449 95511 8889 23017 99015 3336 96617 86426 14224 80344 21574 61529 66580 7799 1264 14323 9425 42811 32562 77844 10844 66222 78172 99070 2150 17659 73943 21320 40020 92675 40641 60313 41101 50741 7926 45148 50892 47981 x", "47981 50892 45148 7926 50741 41101 60313 40641 92675 40020 21320 73943 17659 2150 99070 78172 66222 10844 77844 32562 42811 9425 14323 1264 7799 66580 61529 21574 80344 14224 86426 96617 3336 99015 23017 8889 95511 51449 24621 78506 42742 38489 39091 33419 52447 82456 20377 60988 29835 8003 65445 86751 297 54904 65246 27028 25809 31961 68095 77884 94473 74877 98564 67544 98828 72309 86444 40942 2090 53012 36455 4370 34046 78614"},
	// {"16488 75089 22623 17905 31008 46938 72685 9875 6963 37924 25084 24247 13091 38164 73561 9812 59378 44255 45470 71401 7466 35089 51988 624 78680 95256 30600 62339 26284 88447 55700 42070 44537 25019 56073 72218 27775 25980 3423 31286 49637 24613 83642 6309 1105 14009 x", "14009 1105 6309 83642 24613 49637 31286 3423 25980 27775 72218 56073 25019 44537 42070 55700 88447 26284 62339 30600 95256 78680 624 51988 35089 7466 71401 45470 44255 59378 9812 73561 38164 13091 24247 25084 37924 6963 9875 72685 46938 31008 17905 22623 75089 16488"},
	// {"68754 83534 85 97858 41770 16904 60076 75427 75472 65897 92654 10733 59338 14349 98910 49579 25752 31238 6440 72256 68733 58139 95287 97541 87747 90223 11183 58320 83311 17195 45777 4748 40744 26943 88947 88732 48563 78833 93333 88533 25953 5884 84154 43299 56672 9671 71882 73857 3220 54237 x", "54237 3220 73857 71882 9671 56672 43299 84154 5884 25953 88533 93333 78833 48563 88732 88947 26943 40744 4748 45777 17195 83311 58320 11183 90223 87747 97541 95287 58139 68733 72256 6440 31238 25752 49579 98910 14349 59338 10733 92654 65897 75472 75427 60076 16904 41770 97858 85 83534 68754"},
	// {"64198 41131 94975 58003 27456 71820 2251 40767 70569 24174 14250 37270 37506 58835 99011 42618 54968 75663 95250 90263 79059 22907 52878 42894 81412 58045 60021 12487 87228 12264 37405 71734 70919 33961 34308 41261 82797 80039 68358 48777 24715 1163 52854 89684 60720 26831 97899 8196 13943 22669 35114 56516 34921 68943 66892 81407 54560 90445 5882 60857 32018 80477 58093 11166 42697 30075 81729 45386 88319 51061 34062 12682 55651 34504 13157 69309 891 55579 35712 59324 86120 34678 83087 81441 89747 40388 13294 85694 69350 7655 16107 2465 98504 x", "98504 2465 16107 7655 69350 85694 13294 40388 89747 81441 83087 34678 86120 59324 35712 55579 891 69309 13157 34504 55651 12682 34062 51061 88319 45386 81729 30075 42697 11166 58093 80477 32018 60857 5882 90445 54560 81407 66892 68943 34921 56516 35114 22669 13943 8196 97899 26831 60720 89684 52854 1163 24715 48777 68358 80039 82797 41261 34308 33961 70919 71734 37405 12264 87228 12487 60021 58045 81412 42894 52878 22907 79059 90263 95250 75663 54968 42618 99011 58835 37506 37270 14250 24174 70569 40767 2251 71820 27456 58003 94975 41131 64198"},
	{"Masuk 15 Masuk 16 Masuk 10 Masuk 17 Masuk 16 Masuk 16 Keluar 2 Masuk 17 Keluar 3 x", "5"},
	{"Masuk 10 Masuk 14 Masuk 15 Keluar 1 Masuk 12 Masuk 15 Keluar 2 Masuk 17 Masuk 19 Masuk 13 Masuk 14 Keluar 3 x", "6"},
	{"Masuk 14 Masuk 16 Masuk 13 Masuk 15 Masuk 15 Masuk 15 Masuk 15 Keluar 3 Keluar 3 x", "6"},
	{"Keluar 1 Masuk 16 Masuk 17 Masuk 18 Masuk 19 Masuk 20 Masuk 21 Masuk 22 Masuk 23 Keluar 3 Keluar 3 Keluar 2 x", "8"},
	{"Keluar 1 Masuk 13 Masuk 13 Masuk 12 Masuk 15 Masuk 16 Keluar 2 Masuk 17 Keluar 1 x", "3"},
	{"Keluar 1 Masuk 14 Masuk 15 Masuk 16 Masuk 17 Masuk 18 Masuk 19 Masuk 20 Keluar 3 Masuk 21 Masuk 22 Masuk 23 Masuk 24 Keluar 3 Keluar 1 Keluar 3 Masuk 23 Masuk 22 Keluar 3 x", "13"},
	{"Keluar 1 Masuk 20 Masuk 19 Masuk 18 Masuk 17 Masuk 16 Masuk 15 Masuk 14 Masuk 13 Keluar 3 Keluar 3 Keluar 2 x", "7"},
	{"Masuk 13 Masuk 13 Masuk 12 Keluar 1 Masuk 14 Masuk 15 Masuk 15 Keluar 3 Masuk 20 Masuk 19 Masuk 18 Masuk 17 Masuk 16 Masuk 15 Masuk 14 Masuk 13 Keluar 3 Keluar 3 Keluar 2 Masuk 10 Masuk 14 Masuk 15 Keluar 1 Masuk 12 Masuk 15 Keluar 2 Masuk 17 Masuk 19 Masuk 13 Masuk 14 Keluar 3 Masuk 14 Masuk 15 Masuk 16 Masuk 17 Masuk 18 Masuk 19 Masuk 20 Keluar 3 Masuk 21 Masuk 22 Masuk 23 Masuk 24 Keluar 3 Keluar 1 Keluar 3 Masuk 23 Masuk 22 Keluar 3 Masuk 13 Masuk 13 Masuk 12 Masuk 15 Masuk 16 Keluar 2 Masuk 17 Keluar 1 x", "32"},
	{"Keluar 1 Masuk 14 Masuk 15 Masuk 16 Keluar 3 Masuk 14 Masuk 14 Masuk 14 Masuk 10 Keluar 4 Masuk 13 Masuk 14 Masuk 14 Keluar 2 x", "8"},
	{"Masuk 15 Masuk 16 Masuk 17 Keluar 3 x", "3"},
}
var T4 = TestCase{
	// {"{{{{}}(()})}{{[}]}()", "BELUM SEIMBANG"},
	// {"{}{}{}{}{}()()()()()[][][][][][]", "SUDAH SEIMBANG"},
	// {"{{{{{{{{{{[[[[[[[[[[(((((((((())))))))))]]]]]}}}}}", "BELUM SEIMBANG"},
	// {"{[[({})]]}{[[({})]]}{[[({})]]}{[[({})]]}{[[({})]]}", "SUDAH SEIMBANG"},
	// {"{{}(){{}(){{}(){{}(){{}(){{}(){{}(){{}(){{}(){{}()", "BELUM SEIMBANG"},
	// {"]]]]]]]]]]]]]]]]]]]][[[[[[[[[[[[[[[[[[[[", "BELUM SEIMBANG"},
	// {"[[[[[[[[[[[[[[[[[[[[]]]]]]]]]]}}}}}}}}}}", "BELUM SEIMBANG"},
	// {"({[({[({[({[({[({[({[({[]})]})]})]})]})]})]})]})", "SUDAH SEIMBANG"},
	// {"[({[({[({[({[({[({[({})]})]})]})]})]})]})][", "BELUM SEIMBANG"},
	// {"[[({()})]][[({()})]][[({()})]][[({()})]][[({()})]]", "SUDAH SEIMBANG"},
	{"5 E G C R A", "A C E G R"},
	{"3 A C B", "A B C"},
	{"10 G H V C N M J U E R", "C E G H J M N R U V"},
	{"4 A R T Y ", "A R T Y"},
	{"6 V H I Y F X", "F H I V X Y"},
	{"12 G J U Y I O P K B N M X", "B G I J K M N O P U X Y"},
	{"15 A S D F G H J K L Z X C V B N", "A B C D F G H J K L N S V X Z"},
	{"18 Q W E R T Y U I O P L K J H G F D S", "D E F G H I J K L O P Q R S T U W Y"},
	{"20 P O I U Y T R E W Q S A D X Z C F V G H", "A C D E F G H I O P Q R S T U V W X Y Z"},
	{"26 Z Y X W V U T S R Q P O N M L K J I H G F E D C B A", "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z"},
}

var T5 = TestCase{
	// {"3\naku\nsayang\nkamu", "aku sayang kamu"},
	// {"5\na\nb\nc\nd\ne", "a b c d e"},
	// {"2\ns\nd", "s d"},
	// {"1\noy", "oy"},
	// {"4\nini\ncontoh\nempat\nkalimat", "ini contoh empat kalimat"},
	// {"6\nsatu\ndua\ntiga\nempat\nlima\nenam", "satu dua tiga empat lima enam"},
	// {"7\ndisini\naku\nsendiri\nngga\nsama\nkamu\noy", "disini aku sendiri ngga sama kamu oy"},
	// {"9\nloh\nkok\ngitu\nsih\nloh\nkok\nmarah\njangan\nmarah", "loh kok gitu sih loh kok marah jangan marah"},
	// {"10\na\nc\ns\nv\nb\nd\ns\nf\ns\nd", "a c s v b d s f s d"},
	// {"8\na\nw\no\nk\na\nw\no\nk", "a w o k a w o k"},
	{"3\ngiyuu\ntengen\nkanroji\ntanjiro", "=== giyuu masuk antrian ===\n=== tengen masuk antrian ===\n=== kanroji masuk antrian ===\n=== Maaf antrian penuh ===\n1. giyuu\n2. tengen\n3. kanroji\ngiyuu telah keluar antrian\n1. tengen\n2. kanroji\n3. \nMenambahkan tanjiro ke dalam antrian...\n=== tanjiro masuk antrian ===\n1. tengen\n2. kanroji\n3. tanjiro\n"},
	{"3\naufa\nbimo\ntian\nazizi", "=== aufa masuk antrian ===\n=== bimo masuk antrian ===\n=== tian masuk antrian ===\n=== Maaf antrian penuh ===\n1. aufa\n2. bimo\n3. tian\naufa telah keluar antrian\n1. bimo\n2. tian\n3. \nMenambahkan azizi ke dalam antrian...\n=== azizi masuk antrian ===\n1. bimo\n2. tian\n3. azizi\n"},
	{"3\nzaki\nchris\nskrull\naziz", "=== zaki masuk antrian ===\n=== chris masuk antrian ===\n=== skrull masuk antrian ===\n=== Maaf antrian penuh ===\n1. zaki\n2. chris\n3. skrull\nzaki telah keluar antrian\n1. chris\n2. skrull\n3. \nMenambahkan aziz ke dalam antrian...\n=== aziz masuk antrian ===\n1. chris\n2. skrull\n3. aziz\n"},
	{"3\nmaul\nadi\nkeenan\nalvin", "=== maul masuk antrian ===\n=== adi masuk antrian ===\n=== keenan masuk antrian ===\n=== Maaf antrian penuh ===\n1. maul\n2. adi\n3. keenan\nmaul telah keluar antrian\n1. adi\n2. keenan\n3. \nMenambahkan alvin ke dalam antrian...\n=== alvin masuk antrian ===\n1. adi\n2. keenan\n3. alvin\n"},
	{"3\nsadasd\nfkgjfk\noweiwe\nqwerty", "=== sadasd masuk antrian ===\n=== fkgjfk masuk antrian ===\n=== oweiwe masuk antrian ===\n=== Maaf antrian penuh ===\n1. sadasd\n2. fkgjfk\n3. oweiwe\nsadasd telah keluar antrian\n1. fkgjfk\n2. oweiwe\n3. \nMenambahkan qwerty ke dalam antrian...\n=== qwerty masuk antrian ===\n1. fkgjfk\n2. oweiwe\n3. qwerty\n"},
	{"3\nkoko\nkiki\nkeke\nkaka", "=== koko masuk antrian ===\n=== kiki masuk antrian ===\n=== keke masuk antrian ===\n=== Maaf antrian penuh ===\n1. koko\n2. kiki\n3. keke\nkoko telah keluar antrian\n1. kiki\n2. keke\n3. \nMenambahkan kaka ke dalam antrian...\n=== kaka masuk antrian ===\n1. kiki\n2. keke\n3. kaka\n"},
	{"3\ndada\ndidi\ndudu\ndede", "=== dada masuk antrian ===\n=== didi masuk antrian ===\n=== dudu masuk antrian ===\n=== Maaf antrian penuh ===\n1. dada\n2. didi\n3. dudu\ndada telah keluar antrian\n1. didi\n2. dudu\n3. \nMenambahkan dede ke dalam antrian...\n=== dede masuk antrian ===\n1. didi\n2. dudu\n3. dede\n"},
	{"3\nsatu\ndua\ntiga\nempat", "=== satu masuk antrian ===\n=== dua masuk antrian ===\n=== tiga masuk antrian ===\n=== Maaf antrian penuh ===\n1. satu\n2. dua\n3. tiga\nsatu telah keluar antrian\n1. dua\n2. tiga\n3. \nMenambahkan empat ke dalam antrian...\n=== empat masuk antrian ===\n1. dua\n2. tiga\n3. empat\n"},
	{"3\none\ntwo\nthree\nfour", "=== one masuk antrian ===\n=== two masuk antrian ===\n=== three masuk antrian ===\n=== Maaf antrian penuh ===\n1. one\n2. two\n3. three\none telah keluar antrian\n1. two\n2. three\n3. \nMenambahkan four ke dalam antrian...\n=== four masuk antrian ===\n1. two\n2. three\n3. four\n"},
	{"3\na\nb\nc\nd", "=== a masuk antrian ===\n=== b masuk antrian ===\n=== c masuk antrian ===\n=== Maaf antrian penuh ===\n1. a\n2. b\n3. c\na telah keluar antrian\n1. b\n2. c\n3. \nMenambahkan d ke dalam antrian...\n=== d masuk antrian ===\n1. b\n2. c\n3. d\n"},
}

var T6 = TestCase{
	// {"1\n2\n3\n4\n5", "5\n4\n3\n2\n1\nPaling atas : 5\nBerhasil Mengeluarkan 5\nPaling atas : 4"},
	// {"Apel\nJeruk\nMangga\nPir\nDaging", "Daging\nPir\nMangga\nJeruk\nApel\nPaling atas : Daging\nBerhasil Mengeluarkan Daging\nPaling atas : Pir"},
	// {"a\nb\nc\nd\ne", "e\nd\nc\nb\na\nPaling atas : e\nBerhasil Mengeluarkan e\nPaling atas : d"},
	// {"5\n4\n3\n2\n1", "1\n2\n3\n4\n5\nPaling atas : 1\nBerhasil Mengeluarkan 1\nPaling atas : 2"},
	// {"nastar\nbiskuit\npermen\nkue\nsushi", "sushi\nkue\npermen\nbiskuit\nnastar\nPaling atas : sushi\nBerhasil Mengeluarkan sushi\nPaling atas : kue"},
	// {"pisau\ngarpu\nsendok\nsumpit\ncangkul", "cangkul\nsumpit\nsendok\ngarpu\npisau\nPaling atas : cangkul\nBerhasil Mengeluarkan cangkul\nPaling atas : sumpit"},
	// {"ha\nhi\nhu\nhe\nhm", "hm\nhe\nhu\nhi\nha\nPaling atas : hm\nBerhasil Mengeluarkan hm\nPaling atas : he"},
	// {"lada\nmerica\ngula\ngaram\nmesiu", "mesiu\ngaram\ngula\nmerica\nlada\nPaling atas : mesiu\nBerhasil Mengeluarkan mesiu\nPaling atas : garam"},
	// {"tomat\ntimun\nseledri\nkemangi\napel", "apel\nkemangi\nseledri\ntimun\ntomat\nPaling atas : apel\nBerhasil Mengeluarkan apel\nPaling atas : kemangi"},
	// {"Mangga\nDurian\nLeci\nAnggur\nKompor", "Kompor\nAnggur\nLeci\nDurian\nMangga\nPaling atas : Kompor\nBerhasil Mengeluarkan Kompor\nPaling atas : Anggur"},
	{"5\n2 5", "Hasil Faktorial : 120\nHasil Pangkat : 32"},
	{"6\n7 2", "Hasil Faktorial : 720\nHasil Pangkat : 49"},
	{"7\n9 2", "Hasil Faktorial : 5040\nHasil Pangkat : 81"},
	{"2\n6 3", "Hasil Faktorial : 2\nHasil Pangkat : 216"},
	{"3\n5 3", "Hasil Faktorial : 6\nHasil Pangkat : 125"},
	{"8\n4 3", "Hasil Faktorial : 40320\nHasil Pangkat : 64"},
	{"0\n2 8", "Hasil Faktorial : 1\nHasil Pangkat : 256"},
	{"1\n2 1", "Hasil Faktorial : 1\nHasil Pangkat : 2"},
	{"9\n3 2", "Hasil Faktorial : 362880\nHasil Pangkat : 9"},
	{"4\n2 5", "Hasil Faktorial : 24\nHasil Pangkat : 32"},
}

var T7 = TestCase{
	// {"8\n10 2 3 4 6 7 20 18", "18 20 7 6 4 3 2 10"},
	// {"10\n1 2 3 4 5 6 7 8 9 10", "10 9 8 7 6 5 4 3 2 1"},
	// {"1\n5", "5"},
	// {"3\n8 9 20", "20 9 8"},
	// {"2\n2 9", "9 2"},
	// {"5\n5 10 15 20 25", "25 20 15 10 5"},
	// {"7\n10 20 30 40 50 60 70", "70 60 50 40 30 20 10"},
	// {"6\n6 12 18 24 30 36", "36 30 24 18 12 6"},
	// {"4\n4 8 12 16", "16 12 8 4"},
	// {"9\n9 8 7 6 5 4 3 2 1", "1 2 3 4 5 6 7 8 9"},
	{"2 add 10 2 add 5 6", "8\n10 10 5 5 5 5 5 5"},
	{"5 add 7 2 add 5 3 del 3 add 2 1 del 3", "0"},
	{"3 add 5 6 del 1 del 2", "3\n5 5 5"},
	{"3 add 2 5 add 3 1 add 4 2", "8\n2 2 2 2 2 3 4 4"},
	{"2 add 4 3 add 5 5", "8\n4 4 4 5 5 5 5 5"},
	{"1 add 1 1", "1\n1"},
	{"3 add 1 1 add 2 2 del 1", "2\n2 2"},
	{"2 add 2 3 add 4 3", "6\n2 2 2 4 4 4"},
	{"1 add 2 2", "2\n2 2"},
	{"3 add 4 3 add 2 3 del 2", "4\n4 2 2 2"},
}

var T8 = TestCase{
	// {"6\nGTA DMC Bully Mario Luigi Warcraft\nBully", "Ada"},
	// {"3\nRE4 Dota ML\nBattlefront", "Tidak Ada"},
	// {"1\nWOW\nWOW", "Ada"},
	// {"10\nRE4 Dota ML Bully A B C D Z F\nDota", "Ada"},
	// {"4\nGTA1 GTA2 GTA3 GTA4\nGTA5", "Tidak Ada"},
	// {"2\nRE2 RE3\nRE2", "Ada"},
	// {"3\nRE4 Dota ML\nML", "Ada"},
	// {"6\nRE4 RE2 GTA2 Dota ML GTA5\nRE2", "Ada"},
	// {"5\nRE2 RE4 Dota ML LOL\nWild", "Tidak Ada"},
	// {"4\nRE4 Warcraft Dota ML\nStarwars", "Tidak Ada"},
	{"3 3 4 5", "2\n3\n5"},
	{"2 30 15", "832040\n610"},
	{"3 1 2 3", "1\n1\n2"},
	{"5 9 8 7 4 5", "34\n21\n13\n3\n5"},
	{"2 20 21", "6765\n10946"},
	{"1 25", "75025"},
	{"4 10 12 14 16", "55\n144\n377\n987"},
	{"2 7 15", "13\n610"},
	{"1 28", "317811"},
	{"3 8 7 6", "21\n13\n8"},
	{"4 38 40 41 42", "39088169\n102334155\n165580141\n267914296"},
	{"2 43 40", "433494437\n102334155"},
	{"1 45", "1134903170"},
	{"2 45 45", "1134903170\n1134903170"},
	{"5 45 44 43 42 41", "1134903170\n701408733\n433494437\n267914296\n165580141"},
}