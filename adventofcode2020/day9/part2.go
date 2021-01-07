package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(RawInput, "\n")
	preambleLen := 25

	data := make([]int, len(input))
	for i, v := range input {
		val, _ := strconv.Atoi(v)
		data[i] = val
	}

	//preamble := data[0:preambleLen]

	invalidNum := 0
	for i := preambleLen; i < len(data); i++ {
		section := data[i-preambleLen : i]
		if !isMatch(section, data[i]) {
			fmt.Printf("%v %d\n", section, data[i])
			invalidNum = data[i]
			break
		}
	}

	for i := 0; i < len(data); i++ {
		sum := data[i]
		min := data[i]
		max := data[i]
		for j := i + 1; j < len(data); j++ {
			sum += data[j]
			if data[j] > max {
				max = data[j]
			}
			if data[j] < min {
				min = data[j]
			}
			if sum == invalidNum {
				fmt.Printf("Found it %d + %d = %d", min, max, min+max)
				return
			} else if sum > invalidNum {
				break
			}
		}
	}

	//fmt.Println(preamble)

}

func isMatch(section []int, target int) bool {
	for _, x := range section {
		for _, y := range section {
			if x != y && (x+y) == target {
				return true
			}
		}
	}
	return false
}

const DemoInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

const RawInput = `18
19
46
14
29
45
40
47
25
43
36
22
21
4
32
33
37
38
26
2
42
15
5
13
31
9
6
30
7
14
10
8
69
11
12
25
16
17
22
19
18
20
51
21
24
23
26
15
60
13
29
27
44
42
28
38
30
36
31
32
33
34
35
37
55
39
40
70
41
43
78
58
73
101
60
56
57
75
59
67
71
61
68
110
69
80
72
112
76
79
144
129
114
84
99
113
132
115
148
116
118
120
128
126
187
130
190
152
229
141
151
155
299
209
163
239
183
197
331
212
228
270
231
234
236
244
246
355
256
392
271
342
292
360
296
306
318
346
602
375
380
395
534
456
440
459
470
502
478
480
490
517
527
548
563
1082
1136
896
1044
614
693
698
988
755
1065
1352
835
899
915
942
929
1173
958
1043
970
1007
1075
1090
1529
1750
1663
1307
1312
1533
1369
1391
1705
1590
1654
1734
1764
1805
2923
2649
1936
1899
2633
1928
2681
1977
3956
2165
2981
2619
3926
2676
2840
2703
4292
4547
4383
4628
3244
3388
3539
6659
6602
5143
3827
5316
8558
3905
5322
4142
4817
4784
9069
5379
11206
5516
7536
10701
6242
7293
6632
15190
11260
7366
7732
7681
7969
8959
9521
8611
8047
10139
8689
16605
8926
10300
10163
10895
15262
12148
13925
12874
15025
13974
19229
14313
15047
15098
15335
15413
16370
16016
16658
27899
16736
18852
35245
17615
19089
19821
29984
23769
23043
30309
31705
26799
26848
35105
28287
30382
42183
30748
30433
31351
31429
32386
32674
33394
44635
34351
50281
36704
60293
69456
42864
49891
46812
49842
53647
55086
85047
119298
60961
75099
62099
73297
94916
61784
62780
63815
75250
96654
67745
104928
71055
111242
133358
89676
92706
151460
134938
101898
158753
108733
122831
122745
123060
293691
126595
165971
142995
125599
130525
146305
194115
402424
179788
138800
226064
160731
182382
235256
191574
194604
210631
224643
269464
231478
231564
245576
253585
248659
252194
276830
256124
411266
371362
299531
354846
654618
340519
330374
459212
343113
695365
431041
386178
402205
405235
435274
456121
633683
504783
477140
586095
701736
500853
508318
654377
555655
640050
741024
642644
673487
1288060
802325
716552
729291
1497690
1073685
1273733
821452
807440
891395
912414
933261
977993
1117190
1140903
1210054
1009171
1143497
1063973
1445843
1195705
1282694
2428963
1316131
1390039
1518877
2095183
2337706
1536731
1628892
1698835
2076758
1712847
1803809
1824656
1845675
2962865
1987164
2073144
2150074
2152668
2204876
2207470
2981529
2478399
3128369
2908916
4192040
2706170
2926770
3055608
3790973
3165623
3235566
3327727
3502644
5835686
3516656
3628465
3670331
7131109
5941736
5116386
4278020
4913640
4412346
5186405
4685869
7293617
5834539
5615086
6383335
5632940
9001309
5982378
6221231
6401189
6493350
6563293
6830371
8188513
7145121
10907100
10354082
7948351
11998421
8690366
9528732
8963889
15987022
10318809
9872274
13376235
16136864
11248026
12016275
14983687
17469257
12203609
12383567
26433146
12894539
13056643
13393664
13975492
15093472
26270774
16638717
16912240
18492621
17654255
22066601
25440210
18836163
20191083
21120300
23264301
26870031
33085622
23451635
24399842
24587176
25098148
51578243
35095792
34176943
37980840
35460265
27369156
29068964
37160073
33550957
36146876
62568016
47164749
41311383
44778259
45631293
39027246
52333265
55938995
48549783
50321666
77431413
47851477
48987018
65898559
52467304
83311625
68646749
74487511
69011222
80925135
76187319
62619921
115214565
167547830
136864130
91943008
131266670
80338629
83805505
86878723
87577029
100184742
117561005
139346027
116498226
96838495
129912153
101454322
115087225
213051449
150674830
246410379
154562929
216668887
138807240
195425854
142958550
164144134
267732572
167217352
167915658
170684228
185259827
171382534
244412872
184415524
265598456
198292817
211925720
213336721
334233094
216541547
240261562
253894465
369675351
281765790
335133010
323222764
505860018
328218377
354884270
307102684
331361486
337901580
379143072
338599886
342066762
456803109
355798058
382708341
594634061
410218537
536559485
425262441
429878268
547903033
523644231
535660255
585255951
753480818
588868474
630325448
662900742
788164595
635321061
638464170
645004264
669263066
734941130
680666648
694397944
1055587889
1363661010
738506399
792926878
945878792
1277125049
1108900182
1357298686
1298221803
776203571
1275329712
1124528729
1174124425
1343567390
1219193922
1310992096
1759849790
1527868008
1428247939
1283468434
1314267330
1349929714
1375064592
1917455607
2517415725
1531433277
1514709970
2328094104
2959681216
1722082363
2811336442
1885103753
1900732300
2126133285
1950327996
2884378519
2552776668
2549189017
2502662356
2530186018
2594460530
2633398148
2597735764
2658533026
5542911545
2664197044
2724994306
4032125695
3046143247
4048849002
3236792333
3399813723
3607186116
6636606056
3622814663
6048128775
6109848472
3851060296
4076461281
7094992249
5032848374
5051851373
5079375035
5097122886
5124646548
5192196294
5231133912
6843978449
5322730070
5771137553
7268918028
7448662725
6445956970
7122604528
6859606996
7006999839
7022628386
10068771633
14029628225
7473874959
10514926364
13866606835
7927521577
9109309655
10084699747
12705008871
12219727414
10176497921
10316842842
10355780460
12253762298
10553863982
19522680326
21340481794
12630744549
13305563966
13452956809
13468585356
13882235382
14480874798
14934521416
17199126307
19663173637
15401396536
16583184614
17036831232
23021851713
18104019498
32584894296
24022449338
37475406147
30076544308
20672623302
37709454534
20909644442
22807626280
23184608531
25936308515
26083701358
27187799348
28240085382
26921542165
28403106772
28363110180
29415396214
46507126270
48026283817
35140850730
40058682945
41059280570
43480249582
43694475015
38776642800
49268309889
46993345800
41582267744
43717270722
55590906120
49272754622
44094252973
71720334964
49120917046
92815392061
53005243523
54109341513
91506533399
55284652345
79835923370
104405569391
76408742014
91720758832
73917493530
78835325745
85769988600
80358910544
140841675878
124053385559
113302602708
120126012736
85299538466
123190248152
142088146683
125529659060
93367007595
236492850860
107114585036
102126160569
108289895868
130518083527
174235354249
165424026929
140584190811
170556084577
150326235544
161708280480
192137928453
180961486314
205425551202
166128899144
165658449010
178666546061
200481592631
233951198406
307551711771
187425699035
306713089955
201656903463
195493168164
209240745605
366092245096
238807979395
210416056437
366610491775
302292471291
290910426355
374535389217
311140275388
312034516024
331787348154
327366729490
402138496094
344324995071
344795445205
353084148045
387907291666
374159714225
491392018986
434301147559
382918867199
389082602498
503949374754
397150071627
715274021156
419656802042
449224035832
512708527728
612554552531
593202897646
776673885311
685675664605
623174791412
638507004878
736003015244
659154077644
886868241953
733878047703
763981797113
1269787109152
750234219672
763242316723
1248711542972
780068938826
1048236680142
1862990006798
1418575943704
1620746289656
2003665156855
2198644882530
868880837874
2005790124396
1205757450177
1216377689058
1297661082522
1324182669483
3868780131194
1261681796290
1372385052581
1393032125347
1409388297316
1497120364426
1484112267375
1513476536395
1530303158498
1543311255549
1632123154597
1648949776700
1828305618968
2532348947517
2241265890455
2946610859222
2074638288051
2193063507357
2517830614574
2085258526932
2422135139235
2585864465773
2514038771580
2654713921637
2893804950887
2634066848871
2745794063665
2765417177928
4424108109385
2893500564691
2981232631801
7009972575158
3043779694893
3073614414047
3175434410146
6249048824193
4606987235568
4250440758203
4827130356228
4159896814983
7588219867369
4599297298512
7672911712559
4507393666167
10418705776224
5076849060872
5099903237353
5689473181726
5527871799758
5399484026799
14682884287717
5511211241593
5658917742619
7497722523432
5874733196492
6025012326694
6117394108940
7425875168349
7233511229030
7335331225129
8410337573186
9327289819075
10127169098270
8987027171211
12734815251928
19190277953884
9906877692966
22180606811149
9584242727039
12825359195148
10176752298225
15596350874692
10910695268392
11170128984212
11058401769418
11385944438085
11533650939111
11683930069313
14435349899880
13210064421621
28298000920089
13350905337970
21660820037381
16560801048105
16919573952168
20713234257160
18314316990286
18571269898250
24120759690013
19491120420005
19760995025264
20083629991191
20494937995431
20642644496457
29372718759704
21235154067643
31146939463349
31469574429276
39831651859200
31553339764849
23069874507398
26560969759591
45904924329156
50681729449604
32842025757975
30270479290138
31665222328256
33480375000273
34875118038391
55785982018269
40255933020695
36885586888536
38062390318255
39252115445269
39844625016455
40403639521721
47055907755022
67698552251479
41877798564100
53340353797536
59297544385898
49630844266989
54539448936674
62321989952667
54623214272247
83941494643558
72686650774430
63112505048113
113794234497717
123786119660013
68355493038664
88215471835927
98549659831167
71760704926927`
