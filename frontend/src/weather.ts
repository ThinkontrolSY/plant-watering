//https://worldweather.wmo.int/zh/wxicons.html
export const weatherIcon = (t: string): string => {
  switch (t) {
    case "晴":
      return "/assets/weather/24.png";
    case "少云":
      return "/assets/weather/22.png";
    case "晴间多云":
      return "/assets/weather/21.png";

    case "多云":
      return "/assets/weather/23.png";

    case "阴":
      return "/assets/weather/20.png";
    case "有风":
      return "/assets/weather/26.png";

    case "平静":
      return "/assets/weather/17.png";

    case "微风":
      return "/assets/weather/17.png";
    case "和风":
      return "/assets/weather/17.png";
    case "清风":
      return "/assets/weather/17.png";
    case "强风/劲风":
      return "/assets/weather/26.png";
    case "疾风":
      return "/assets/weather/26.png";
    case "大风":
      return "/assets/weather/26.png";
    case "烈风":
      return "/assets/weather/26.png";
    case "风暴":
      return "/assets/weather/2.png";
    case "狂爆风":
      return "/assets/weather/2.png";
    case "飓风":
      return "/assets/weather/26.png";
    case "热带风暴":
      return "/assets/weather/2.png";
    case "霾":
      return "/assets/weather/19.png";
    case "中度霾":
      return "/assets/weather/19.png";
    case "重度霾":
      return "/assets/weather/19.png";
    case "严重霾":
      return "/assets/weather/19.png";
    case "阵雨":
      return "/assets/weather/12.png";
    case "雷阵雨":
      return "/assets/weather/2.png";

    case "雷阵雨并伴有冰雹":
      return "/assets/weather/3.png";

    case "小雨":
      return "/assets/weather/15.png";
    case "中雨":
      return "/assets/weather/14.png";
    case "大雨":
      return "/assets/weather/9.png";
    case "暴雨":
      return "/assets/weather/9.png";
    case "大暴雨":
      return "/assets/weather/9.png";
    case "特大暴雨":
      return "/assets/weather/2.png";
    case "强阵雨":
      return "/assets/weather/9.png";
    case "强雷阵雨":
      return "/assets/weather/2.png";
    case "极端降雨":
      return "/assets/weather/2.png";
    case "毛毛雨/细雨":
      return "/assets/weather/15.png";
    case "雨":
      return "/assets/weather/14.png";
    case "小雨-中雨":
      return "/assets/weather/12.png";
    case "中雨-大雨":
      return "/assets/weather/12.png";
    case "大雨-暴雨":
      return "/assets/weather/12.png";
    case "暴雨-大暴雨":
      return "/assets/weather/12.png";
    case "大暴雨-特大暴雨":
      return "/assets/weather/12.png";
    case "雨雪天气":
      return "/assets/weather/8.png";
    case "雨夹雪":
      return "/assets/weather/8.png";
    case "阵雨夹雪":
      return "/assets/weather/8.png";

    case "冻雨":
      return "/assets/weather/13.png";
    case "雪":
      return "/assets/weather/6.png";

    case "阵雪":
      return "/assets/weather/5.png";
    case "小雪":
      return "/assets/weather/7.png";

    case "中雪":
      return "/assets/weather/6.png";
    case "大雪":
      return "/assets/weather/6.png";
    case "暴雪":
      return "/assets/weather/4.png";
    case "小雪-中雪":
      return "/assets/weather/7.png";
    case "中雪-大雪":
      return "/assets/weather/6.png";
    case "大雪-暴雪":
      return "/assets/weather/4.png";

    case "浮尘":
      return "/assets/weather/1.png";
    case "扬沙":
      return "/assets/weather/1.png";
    case "沙尘暴":
      return "/assets/weather/1.png";
    case "强沙尘暴":
      return "/assets/weather/1.png";
    case "龙卷风":
      return "/assets/weather/35.png";
    case "雾":
      return "/assets/weather/16.png";

    case "浓雾":
      return "/assets/weather/16.png";
    case "强浓雾":
      return "/assets/weather/16.png";
    case "轻雾":
      return "/assets/weather/17.png";
    case "大雾":
      return "/assets/weather/16.png";
    case "特强浓雾":
      return "/assets/weather/16.png";
    case "热":
      return "/assets/weather/31.png";
    case "冷":
      return "/assets/weather/32.png";
    case "未知":
      return "/assets/weather/unknown.png";
  }
  return "/assets/weather/unknown.png";
};
