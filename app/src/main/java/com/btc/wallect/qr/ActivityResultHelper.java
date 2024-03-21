package com.btc.wallect.qr;

import static android.app.Activity.RESULT_OK;

import android.Manifest;
import android.app.Activity;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.graphics.Bitmap;
import android.text.TextUtils;
import android.util.Log;
import android.widget.Toast;

import androidx.core.app.ActivityCompat;
import androidx.core.content.ContextCompat;

import com.btc.wallect.utils.LogUntil;
import com.example.scanzxing.zxing.android.CaptureActivity;
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
    public void goScan(Activity activity) {
        if (ContextCompat.checkSelfPermission(activity, Manifest.permission.CAMERA) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(activity, new String[]{Manifest.permission.CAMERA}, 1);
        } else {
            Intent intent = new Intent(activity, CaptureActivity.class);
            activity.startActivityForResult(intent, Constantes.REQUEST_CODE_SCAN);
        }

    }
}
