package com.btc.wallect.utils;

import android.app.Activity;
import android.text.method.HideReturnsTransformationMethod;
import android.text.method.PasswordTransformationMethod;
import android.widget.EditText;
import android.widget.ImageView;

import com.btc.wallect.R;

public class UiUtils {
    /**
     * 隐藏密码
     *
     * @param activity
     * @param et：输入框控件
     * @param iv：小眼睛控件
     * @param boo：是否隐藏
     */
    public static void setHidePassword(Activity activity, EditText et, ImageView iv, boolean boo) {
        if (activity != null) {
            activity.runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    if (et != null && iv != null) {
                        if (boo) {
                            //隐藏输入框内容
                            et.setTransformationMethod(PasswordTransformationMethod.getInstance());
                            //改变小眼睛控件UI
                            iv.setBackgroundResource(R.mipmap.img_close_eyes);

                        } else {
                            //显示输入框内容
                            et.setTransformationMethod(HideReturnsTransformationMethod.getInstance());
                            iv.setBackgroundResource(R.mipmap.img_open_eyes);
                        }
                    }
                }
            });
        }
    }


}
