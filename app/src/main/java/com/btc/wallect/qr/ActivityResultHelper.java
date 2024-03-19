package com.btc.wallect.qr;

import static android.app.Activity.RESULT_OK;

import android.content.Intent;
import android.graphics.Bitmap;
import android.text.TextUtils;
import android.util.Log;
import android.widget.Toast;

import com.btc.wallect.utils.LogUntil;
import com.example.scanzxing.zxing.common.Constantes;

public class ActivityResultHelper {
    public interface OnActivityResultListener {
        void onActivityResult(int requestCode, int resultCode, Intent data);
    }

    private OnActivityResultListener listener;

    public ActivityResultHelper(OnActivityResultListener listener) {
        this.listener = listener;
    }

    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        if (listener != null) {

            listener.onActivityResult(requestCode, resultCode, data);
        }
    }
}
