package com.btc.wallect.utils;

import android.annotation.SuppressLint;
import android.text.TextUtils;
import android.text.format.DateUtils;

import java.io.File;
import java.text.ParseException;
import java.text.ParsePosition;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.TimeZone;

public class DateTimeUtil {
    private volatile static DateTimeUtil dateTimeUtil = null;
    private Date date = new Date();
    @SuppressLint("SimpleDateFormat")
    private final SimpleDateFormat simpleDateFormat = new SimpleDateFormat();

    public static DateTimeUtil getInstance() {
        if (null == dateTimeUtil) {
            synchronized (DateTimeUtil.class) {
                if (null == dateTimeUtil) {
                    dateTimeUtil = new DateTimeUtil();
                }
            }
        }
        return dateTimeUtil;
    }

    /**
     * 获取当前系统时间戳
     *
     * @return
     */
    public long getNowLongTime() {
        return System.currentTimeMillis();
    }

    /**
     * 根据时间毫秒数格式化时间
     *
     * @param time
     * @param style
     * @return
     */
    public String formatDateTime(long time, int style) {
        date.setTime(time);
        switch (style) {
            case 0:
                simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm:ss");
                break;
            case 1:
                simpleDateFormat.applyPattern("yyyy/MM/dd HH:mm:ss");
                break;
            case 2:
                simpleDateFormat.applyPattern("MM-dd HH:mm");
                break;
            case 3:
                simpleDateFormat.applyPattern("MM-dd HH:mm:ss");
                break;
            case 4:
                simpleDateFormat.applyPattern("M-d");
                break;
            case 5:
                simpleDateFormat.applyPattern("M-d HH:mm:ss");
                break;
            case 6:
                simpleDateFormat.applyPattern("mm:ss");
                break;
            case 7:
                simpleDateFormat.applyPattern("yyyy/M/d");
                break;
            case 8:
                simpleDateFormat.applyPattern("yyyy/MM/dd HH:mm");
                break;
            case 9:
                simpleDateFormat.applyPattern("M/d HH:mm");
                break;
            case 10:
                simpleDateFormat.applyPattern("yyyy年MM月");
                break;
            case 11:
                simpleDateFormat.applyPattern("HH:mm:ss");
                break;
            case 12:
                simpleDateFormat.applyPattern("M月d日");
                break;
            case 13:
                simpleDateFormat.applyPattern("HH:mm");
                break;
            case 14:
                simpleDateFormat.applyPattern("yyyy-MM-dd");
                break;
            case 15:
                simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm");
                break;
            case 16:
                simpleDateFormat.applyPattern("M月d日 HH:mm");
                break;
            default:
                simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm:ss");
                break;
        }
        return simpleDateFormat.format(date);
    }


    /**
     * 将2013:10:08 11:48:07如此格式的时间,转化为毫秒数
     *
     * @param datetime 字符串时间
     * @return 毫秒数
     */
    public long dateTimeToMs(String datetime) {
        if (TextUtils.isEmpty(datetime)) {
            return 0;
        }
        try {
            simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm:ss");
            date = simpleDateFormat.parse(datetime);
        } catch (ParseException e) {
            e.printStackTrace();
            return 0;
        }
        return date.getTime();
    }

    /**
     * 将20131008114807如此格式的时间,转化为毫秒数
     *
     * @param datetime 字符串时间
     * @return 毫秒数
     */
    public long dateTime2Ms(String datetime) {
        if (TextUtils.isEmpty(datetime) || datetime.length() != 14) {
            return 0;
        }
        String stringBuilder = datetime.substring(0, 4) +
                ":" +
                datetime.substring(4, 6) +
                ":" +
                datetime.substring(6, 8) +
                " " +
                datetime.substring(8, 10) +
                ":" +
                datetime.substring(10, 12) +
                ":" +
                datetime.substring(12, 14);
        return dateTimeToMs(stringBuilder);
    }

    /**
     * 将2013-10-08如此格式的时间,转化为毫秒数
     *
     * @param datetime 字符串时间
     * @return 毫秒数
     */
    public long dateTime3ToMs(String datetime) {
        if (TextUtils.isEmpty(datetime)) {
            return 0;
        }
        try {
            simpleDateFormat.applyPattern("yyyy-MM-dd");
            date = simpleDateFormat.parse(datetime);
        } catch (ParseException e) {
            e.printStackTrace();
            return 0;
        }
        return date.getTime();
    }

    /**
     * 计算两个日期相隔的天数.
     *
     * @param d1
     * @param d2
     * @return 返回两个日期相隔的天数, 如果是同一天返回0.
     */
    public int getDaysBetween(Calendar d1, Calendar d2) {
        if (d1.after(d2)) {
            java.util.Calendar swap = d1;
            d1 = d2;
            d2 = swap;
        }
        int days = d2.get(java.util.Calendar.DAY_OF_YEAR) - d1.get(java.util.Calendar.DAY_OF_YEAR);
        int y2 = d2.get(java.util.Calendar.YEAR);
        if (d1.get(java.util.Calendar.YEAR) != y2) {
            d1 = (java.util.Calendar) d1.clone();
            do {
                days += d1.getActualMaximum(java.util.Calendar.DAY_OF_YEAR);
                d1.add(java.util.Calendar.YEAR, 1);
            }
            while (d1.get(java.util.Calendar.YEAR) != y2);
        }
        return days;
    }


    /**
     * 根据文件路径,返回文件修改时间
     *
     * @param path
     * @return
     */
    public long getFileDate(String path) {
        File file = new File(path);
        return file.lastModified();
    }


    /**
     * 将毫秒转化为时分秒，用于播放录音显示时长
     *
     * @param timeLength
     * @return
     */
    @SuppressLint("SimpleDateFormat")
    public String ssToTime(long timeLength) {
        SimpleDateFormat formatter = null;
        if (timeLength >= 3600000) {
            formatter = new SimpleDateFormat("HH:mm:ss");
        } else {
            formatter = new SimpleDateFormat("mm:ss");
        }
        formatter.setTimeZone(TimeZone.getTimeZone("GMT+00:00"));
        return formatter.format(timeLength);
    }

    /**
     * 获取当前年月日
     *
     * @return
     */
    public long getNowTime() {

        return System.currentTimeMillis();
    }


    /**
     * 获取今日时间戳(只精确到天)
     *
     * @return
     */
    public long getTodayNowTime() {
        simpleDateFormat.applyPattern("yyyy-MM-dd");
        date.setTime(getNowLongTime());
        String format = simpleDateFormat.format(date);
        try {
            date = simpleDateFormat.parse(format);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return date.getTime();
    }


    /**
     * 获取当前时间的总分钟数
     *
     * @return
     */
    public int getCurrentMinute() {
        // 当前日期
        Calendar cal = Calendar.getInstance();
        // 获取小时
        int hour = cal.get(Calendar.HOUR_OF_DAY);
        // 获取分钟
        int minute = cal.get(Calendar.MINUTE);
        // 从0:00分开是到目前为止的分钟数
        return hour * 60 + minute;
    }


    /**
     * 获取指定时间的总分钟数
     *
     * @param time
     * @return
     */
    public int getCurrentMinute(long time) {
        if (time == 0) {
            return 0;
        }
        // 当前日期
        Calendar cal = Calendar.getInstance();
        //设置时间
        cal.setTime(new Date(time));
        // 获取小时
        int hour = cal.get(Calendar.HOUR_OF_DAY);
        // 获取分钟
        int minute = cal.get(Calendar.MINUTE);
        // 从0:00分开是到目前为止的分钟数
        return hour * 60 + minute;
    }


    /**
     * 获取指定时间的long值（精确到小时）
     *
     * @param time
     * @return
     */
    public long getDateHour(long time) {
        simpleDateFormat.applyPattern("yyyy-MM-dd HH");
        date.setTime(time);
        String format = simpleDateFormat.format(date);
        try {
            date = simpleDateFormat.parse(format);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return date.getTime();
    }


    /**
     * 获取到指定时间的long值（精确到天）
     *
     * @param time
     * @return
     */
    public long getDate(long time) {
        simpleDateFormat.applyPattern("yyyy-MM-dd");
        date.setTime(time);
        String format = simpleDateFormat.format(date);
        try {
            date = simpleDateFormat.parse(format);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return date.getTime();
    }


    /**
     * 获取增加指定天数后系统的时间long（精确到天）
     *
     * @param days
     * @return
     */
    public long getIncreaseTime(int days) {
        //日期
        Calendar c = Calendar.getInstance();
        //日期增加到指定天数后
        c.add(Calendar.DAY_OF_MONTH, days);
        long time = c.getTime().getTime();
        long date = getDate(time);
        return date;
    }


    /**
     * 获取两个时间点，间隔多少分钟
     *
     * @param startTime
     * @param endTime
     * @return
     */
    public int timeInterval(long startTime, long endTime) {
        Date startDate = new Date(startTime);
        Date endDate = new Date(endTime);
        simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm:ss");
        try {
            startDate = simpleDateFormat.parse(simpleDateFormat.format(startDate));
            endDate = simpleDateFormat.parse(simpleDateFormat.format(endDate));
        } catch (ParseException e) {
            e.printStackTrace();
        }

        Calendar calendar = Calendar.getInstance();
        calendar.setTime(startDate);
        long time1 = calendar.getTimeInMillis();


        calendar.setTime(endDate);
        long time2 = calendar.getTimeInMillis();


        long betweenDays = (time2 - time1) / 1000 / 60;
        return Integer.parseInt(String.valueOf(betweenDays));

    }


    /**
     * 判断当前总分钟对应的时间是否处于两个总分钟数对应的时间点之间
     *
     * @param startMinutes
     * @param endMinutes
     * @return
     */
    public boolean judgingTime(int startMinutes, int endMinutes) {
        //获取当前时间的总分钟数
        int currentMinute = getCurrentMinute();
        if (currentMinute >= startMinutes && currentMinute <= endMinutes) {
            return true;
        } else {
            return false;
        }
    }


    /**
     * 将18:00:00如此格式的时间 转化为今日指定时间的时间戳（精确到毫秒）
     *
     * @param time
     * @return
     */
    public long getTodayAppointTime(String time) {
        String dateTime = getNowTime() + " " + time;
        simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm:ss");
        ParsePosition position = new ParsePosition(0);
        date = simpleDateFormat.parse(dateTime, position);
        return date.getTime();
    }


    /**
     * 获取两个时间点，间隔多少分钟
     *
     * @param startTime
     * @param endTime
     * @return
     */
    public int timeIntervalinutes(long startTime, long endTime) {
        Date startDate = new Date(startTime);
        Date endDate = new Date(endTime);
        simpleDateFormat.applyPattern("yyyy-MM-dd HH:mm:ss");
        try {
            startDate = simpleDateFormat.parse(simpleDateFormat.format(startDate));
            endDate = simpleDateFormat.parse(simpleDateFormat.format(endDate));
        } catch (ParseException e) {
            e.printStackTrace();
        }

        Calendar calendar = Calendar.getInstance();
        calendar.setTime(startDate);
        long time1 = calendar.getTimeInMillis();


        calendar.setTime(endDate);
        long time2 = calendar.getTimeInMillis();


        long betweenDays = (time2 - time1) / 1000 / 60;
        return Integer.parseInt(String.valueOf(betweenDays));

    }

    /**
     * 获取今天周几
     *
     * @return
     */
    public String[] getWeek() {
        //week:星期- 星期二 ...
        simpleDateFormat.applyPattern("EEEE");
        String week = simpleDateFormat.format(getNowLongTime());
        String allWeekStr = "【6:星期日】【0:星期一】【1:星期二】【2:星期三】【3:星期四】【4:星期五】【5:星期六】";
//        weekIndex: 1 2 ...
//        Calendar cal = Calendar.getInstance();
//        int weekIndex = cal.get(Calendar.DAY_OF_WEEK) - 1;

        // 6/星期日   0/星期一   1/星期二 2/星期三  3/星期四   4/星期五  5/星期六
        String[] strings = {"6", "0", "1", "2", "3", "4", "5"};
        Calendar cal = Calendar.getInstance();
        int weekIndex = cal.get(Calendar.DAY_OF_WEEK) - 1;
        String str = strings[weekIndex];
        return new String[]{str, week, allWeekStr};
    }

    /**
     * 时长转换
     *
     * @param duration
     * @return
     */
    public String timeCovert(int duration) {
        if (duration <= 0) {
            return "0秒";
        }
        if (duration >= 60) {
            return (duration / 60) + "分" + ((duration % 60) < 10 ? "0" + (duration % 60) : (duration % 60)) + "秒";
        } else {
            return duration + "秒";
        }
    }

    /**
     * 判断时间是否是今日
     *
     * @param time
     * @return
     */
    public boolean isToday(long time) {
        return DateUtils.isToday(time);
    }

    /**
     * 转换为分钟
     *
     * @param time
     * @return
     */
    public int convertToMinutes(String time) {
        simpleDateFormat.applyPattern("HH:mm");
        Date parse = null;
        try {
            parse = simpleDateFormat.parse(time);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return (parse.getHours() * 60) + parse.getMinutes();
    }

    /**
     * 获取N天前的日期或N天后日期(负数之前，正数之后)
     *
     * @param distanceDay 例如：-7/七天前日期   7/七天后日期
     * @return
     */
    public String getOldDate(int distanceDay) {
        Date beginDate = new Date();
        Calendar date = Calendar.getInstance();
        date.setTime(beginDate);
        date.set(Calendar.DATE, date.get(Calendar.DATE) + distanceDay);
        long time = date.getTime().getTime();
        return formatDateTime(time, 14);
    }

    /**
     * 获取N天前的日期时间戳或N天后日期时间戳(负数之前，正数之后)
     *
     * @param distanceDay 例如：-7/七天前日期   7/七天后日期
     * @return
     */
    public long getOldTime(int distanceDay) {
        Date beginDate = new Date();
        Calendar date = Calendar.getInstance();
        date.setTime(beginDate);
        date.set(Calendar.DATE, date.get(Calendar.DATE) + distanceDay);
        return date.getTime().getTime();
    }

    /**
     * 将秒数转为时分秒
     * 00:00
     * 00:00:00
     *
     * @param second
     * @return
     */
    public String getChange(int second) {
        String time;
        if (second < 60) {
            time = String.format("00:%02d", second % 60);
        } else if (second < 3600) {
            time = String.format("%02d:%02d", second / 60, second % 60);
        } else {
            time = String.format("%02d:%02d:%02d", second / 3600, second % 3600 / 60, second % 60);
        }
        return time;
    }

    /**
     * 根据秒数转化为时分秒
     * second + "秒"
     * minute + "分" + second + "秒"
     * hour + "小时" + minute + "分" + second + "秒"
     *
     * @param second
     * @return
     */
    public String getChangeTime(int second) {
        if (second < 60) {
            return second + "秒";
        }
        if (second < 3600) {
            int minute = second / 60;
            second = second - minute * 60;
            return minute + "分" + second + "秒";
        }
        int hour = second / 3600;
        int minute = (second - hour * 3600) / 60;
        second = second - hour * 3600 - minute * 60;
        return hour + "小时" + minute + "分" + second + "秒";
    }

    /**
     * 判断如果是今天返回 "今天"  否则返回"M月d日 HH:mm"
     *
     * @param date
     * @return
     */
    public String getChangeTodayDate(long date) {
        if (date == 0) {
            return "";
        }
        if (DateUtils.isToday(date)) {
            return "今天";
        }
        return formatDateTime(date, 16);
    }

    /**
     * 判断如果是今天返回 "今天"+HH:mm  否则返回"M月d日 HH:mm"
     *
     * @param date
     * @return
     */
    public String getChangeTodayDateTime(long date) {
        if (date == 0) {
            return "";
        }
        if (DateUtils.isToday(date)) {
            return "今天" + formatDateTime(date, 13);
        }
        return formatDateTime(date, 16);
    }

    /**
     * 时间戳转成提示性日期格式：今天、昨天、M月d日
     *
     * @param date
     * @return
     */
    public String getChangeDateTimeFormat(long date) {
        if (date == 0) {
            return "";
        }
        if (DateUtils.isToday(date)) {
            return "今天";
        }
        //现在时间 yyyy-MM-dd
        String today = formatDateTime(getNowLongTime(), 14);
        //昨天 yyyy-MM-dd
        Calendar cal = Calendar.getInstance();
        cal.add(Calendar.DATE, -1);
        String yesterday = formatDateTime(cal.getTime().getTime(), 14);
        if (today.equals(yesterday)) {
            return "昨天";
        }
        return formatDateTime(date, 12);
    }

}
