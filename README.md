# tthead
二地点の緯度経度から距離の算出を行う．  

<img src="./img/tthead_icon.svg" width="100">

## Description
私は地理情報を用いた研究を行っている．その中で二点間の緯度経度から距離の算出を行うことがあった．しかし，距離の算出方法は一つではない．私が行った研究内では，harversine公式を使用し，距離を算出した．  
本ソフトウェアでは二点間の緯度経度から距離を算出するが，harversine公式以外にヒュベニの公式での距離算出を行えるソフトウェアの開発を行う．またosm(Open Street Map)データを想定したファイルの読み込み，距離を算出するファイルを出力するなどの機能を加える．

## Usage
```
edkd [OPTIONS...] [NUMBERs...|FIREs...]  
　OPTIONS:  
　　-H,--Hubeny ヒュベニの公式で距離を算出．デフォルトではharversine公式を使用して算出する．  
　　-x,--same-xcoord 経度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の緯度を引数にする．  
　　-y,--same-ycoord 緯度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の経度を引数にする．  
　　-r,--radian-method 弧度法を引数にする．  
　　-h,--help Usageを表示する．  
    
　ARGUMENTS   
　　NUMBERs... 二地点の緯度経度を引数にする．デフォルトではポイント１経度，ポイント１緯度，ポイント2経度，ポイント2緯度．但しオプション-xの場合，ポイント１緯度，ポイント2緯度，オプション-yの場合，ポイント１経度，ポイント2経度  
　　FIREs... 二地点の緯度経度がまとめられたcsvファイルを引数にする．csvファイルはosmを使う場合都合が良いため．デフォルトでは出力ファイルは"ykgeo_output.csv"，変えたい場合は2つ目の引数で設定する．  
  ```
  
## icon
<img src="./img/tthead_icon.svg" width="100">
