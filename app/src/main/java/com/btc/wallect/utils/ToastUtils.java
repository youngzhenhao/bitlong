package com.btc.wallect.utils;

import android.annotation.SuppressLint;
import android.content.Context;
import android.graphics.Color;
import android.os.Handler;
import android.os.Looper;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.StringRes;

import com.btc.wallect.WallectApp;
import com.btc.wallect.R;

public class ToastUtils {
    private static ToastAdapter mAdapter;
    private static Toast mToast;

    public static void showToast(Context context, String msg) {  //直接显示字符串
        showToast(context, msg, Toast.LENGTH_SHORT);
    }

    public static void showToast(Context context, @StringRes int resId) {
        showToast(context, resId, Toast.LENGTH_SHORT);
    }

    public static void showToast(Context context, @StringRes int resId, int duration) {
        showToast(context, context.getResources().getText(resId), duration);
    }

    public static void showToast(Context context, CharSequence msg, int duration) {
        boolean b = mAdapter == null || mAdapter.displayable();
        if (!b)
            return;
        if (Looper.getMainLooper() == Looper.myLooper())
            obtainAndShowToast(context, msg, duration);
        else
            showToastOnUiThread(context, msg, duration);
    }

    public static void cancelToast() {  //取消
        boolean b = mToast != null && (mAdapter == null || mAdapter.cancellable());
        if (!b)
            return;
        mToast.cancel();
    }

    public interface ToastAdapter {
        boolean displayable();

        boolean cancellable();
    }

    public static void setToastAdapter(ToastAdapter adapter) {
        mAdapter = adapter;
    }

    private static void showToastOnUiThread(final Context context, final CharSequence msg, final int
            duration) {
        new Handler(Looper.getMainLooper()).post(new Runnable() {
            @Override
            public void run() {
                obtainAndShowToast(context, msg, duration);
            }
        });
    }

    @SuppressLint("ShowToast")
    private static void obtainAndShowToast(final Context context, final CharSequence msg, final int
            duration) {
        if (mToast == null) {
            mToast = Toast.makeText(context.getApplicationContext(), msg, duration);
        } else {
            mToast.setText(msg);
            mToast.setDuration(duration);
        }
        mToast.show();
    }


    public static void initToast(String txt,int color) {
        if (mToast == null) {
            mToast = new Toast(WallectApp.getContext());
            View view = LayoutInflater.from(WallectApp.getContext()).inflate(R.layout.toast_custom, null, false);
            TextView tipsText = view.findViewById(R.id.toast_text);
            tipsText.setText(txt);
            if(color==0){
                tipsText.setTextColor(Color.parseColor("#41BF71"));
            } else if (color==1) {
                tipsText.setTextColor(Color.parseColor("#EC3468"));
            }
            mToast.setView(view);
            mToast.show();
        }
    }
}
