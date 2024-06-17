import w24 from "./assets/weather/24.png";
import w26 from "./assets/weather/26.png";
import w17 from "./assets/weather/17.png";
import w19 from "./assets/weather/19.png";
import w12 from "./assets/weather/12.png";
import w2 from "./assets/weather/2.png";
import w3 from "./assets/weather/3.png";
import w15 from "./assets/weather/15.png";
import w14 from "./assets/weather/14.png";
import w9 from "./assets/weather/9.png";
import w8 from "./assets/weather/8.png";
import w6 from "./assets/weather/6.png";
import w5 from "./assets/weather/5.png";
import w7 from "./assets/weather/7.png";
import w4 from "./assets/weather/4.png";
import w1 from "./assets/weather/1.png";
import w35 from "./assets/weather/35.png";
import w16 from "./assets/weather/16.png";
import w31 from "./assets/weather/31.png";
import w32 from "./assets/weather/32.png";
import unknown from "./assets/weather/unknown.png";
import w22 from "./assets/weather/22.png";
import w21 from "./assets/weather/21.png";
import w23 from "./assets/weather/23.png";
import w20 from "./assets/weather/20.png";
import w13 from "./assets/weather/13.png";

//https://worldweather.wmo.int/zh/wxicons.html
export const weatherIcon = (t: string): string => {
  switch (t) {
    case "晴":
      return w24;
    case "少云":
      return w22;
    case "晴间多云":
      return w21;

    case "多云":
      return w23;

    case "阴":
      return w20;
    case "有风":
      return w26;

    case "平静":
      return w17;

    case "微风":
      return w17;
    case "和风":
      return w17;
    case "清风":
      return w17;
    case "强风/劲风":
      return w26;
    case "疾风":
      return w26;
    case "大风":
      return w26;
    case "烈风":
      return w26;
    case "风暴":
      return w2;
    case "狂爆风":
      return w2;
    case "飓风":
      return w26;
    case "热带风暴":
      return w2;
    case "霾":
      return w19;
    case "中度霾":
      return w19;
    case "重度霾":
      return w19;
    case "严重霾":
      return w19;
    case "阵雨":
      return w12;
    case "雷阵雨":
      return w2;

    case "雷阵雨并伴有冰雹":
      return w3;

    case "小雨":
      return w15;
    case "中雨":
      return w14;
    case "大雨":
      return w9;
    case "暴雨":
      return w9;
    case "大暴雨":
      return w9;
    case "特大暴雨":
      return w2;
    case "强阵雨":
      return w9;
    case "强雷阵雨":
      return w2;
    case "极端降雨":
      return w2;
    case "毛毛雨/细雨":
      return w15;
    case "雨":
      return w14;
    case "小雨-中雨":
      return w12;
    case "中雨-大雨":
      return w12;
    case "大雨-暴雨":
      return w12;
    case "暴雨-大暴雨":
      return w12;
    case "大暴雨-特大暴雨":
      return w12;
    case "雨雪天气":
      return w8;
    case "雨夹雪":
      return w8;
    case "阵雨夹雪":
      return w8;

    case "冻雨":
      return w13;
    case "雪":
      return w6;

    case "阵雪":
      return w5;
    case "小雪":
      return w7;

    case "中雪":
      return w6;
    case "大雪":
      return w6;
    case "暴雪":
      return w4;
    case "小雪-中雪":
      return w7;
    case "中雪-大雪":
      return w6;
    case "大雪-暴雪":
      return w4;

    case "浮尘":
      return w1;
    case "扬沙":
      return w1;
    case "沙尘暴":
      return w1;
    case "强沙尘暴":
      return w1;
    case "龙卷风":
      return w35;
    case "雾":
      return w16;

    case "浓雾":
      return w16;
    case "强浓雾":
      return w16;
    case "轻雾":
      return w17;
    case "大雾":
      return w16;
    case "特强浓雾":
      return w16;
    case "热":
      return w31;
    case "冷":
      return w32;
    case "未知":
      return unknown;
  }
  return unknown;
};
