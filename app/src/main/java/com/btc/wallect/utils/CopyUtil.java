package com.btc.wallect.utils;

import android.content.ClipData;
import android.content.ClipboardManager;
import android.content.Context;

import com.btc.wallect.WallectApp;

public class CopyUtil {

    public static void copyClicks(String text) {
        Context context = WallectApp.getContext();
        // 获取系统剪贴板
        ClipboardManager clipboard = (ClipboardManager) context.getSystemService(Context.CLIPBOARD_SERVICE);

        // 创建一个剪贴数据集，包含一个普通文本数据条目（需要复制的数据）
        ClipData clipData = ClipData.newPlainText(null, text);

        // 把数据集设置（复制）到剪贴板
        clipboard.setPrimaryClip(clipData);

      //  ToastUtils.showToast(context, R.string.app_txt43);

    }

    //从剪切板粘贴文本
    public static String pasteClicks() {
        String text = "";
        Context context = WallectApp.getContext();
        // 获取系统剪贴板
        ClipboardManager clipboard = (ClipboardManager) context.getSystemService(Context.CLIPBOARD_SERVICE);
        // 获取剪贴板的剪贴数据集
        ClipData clipData = clipboard.getPrimaryClip();

        if (clipData != null && clipData.getItemCount() > 0) {
            // 从数据集中获取（粘贴）第一条文本数据
            text = clipData.getItemAt(0).getText().toString();
        }
        return text;
    }



}
