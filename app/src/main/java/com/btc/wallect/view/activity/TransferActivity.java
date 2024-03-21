package com.btc.wallect.view.activity;

import static com.btc.wallect.view.activity.fragment.WallectFragment.REQUEST_CODE;

import android.content.Intent;
import android.graphics.Bitmap;
import android.os.Bundle;
import android.text.Editable;
import android.text.TextUtils;
import android.text.TextWatcher;
import android.util.Log;
import android.view.View;
import android.view.WindowManager;
import android.widget.EditText;


import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import com.btc.wallect.R;
import com.btc.wallect.qr.ActivityResultHelper;
import com.btc.wallect.utils.LogUntil;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.example.scanzxing.zxing.android.CaptureActivity;
import com.example.scanzxing.zxing.common.Constantes;

import butterknife.BindView;
import butterknife.OnClick;

public class TransferActivity extends BaseActivity {
    private ActivityResultHelper helper;

    private static final int MAX_AMOUNT = 1000; // 设置允许的最大金额
    @BindView(R.id.ed_text_amount)
    EditText amountEditText;

    @Override
    protected int setContentView() {
        return R.layout.transfer;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle("转账");
        setTitleTxtColor(2);
        setTitleBackgroundColor(1);
        setImgBack(true);
        initScan();
        // 监听金额输入框的文本变化
        amountEditText.addTextChangedListener(new TextWatcher() {
            @Override
            public void beforeTextChanged(CharSequence s, int start, int count, int after) {
                // 在文本改变之前的操作，这里不需要实现
            }

            @Override
            public void onTextChanged(CharSequence s, int start, int before, int count) {
                // 在文本改变时的操作，这里不需要实现
            }

            @Override
            public void afterTextChanged(Editable s) {
                // 在文本改变之后的操作
                if (!s.toString().isEmpty()) {
                    try {
                        int amount = Integer.parseInt(s.toString());
                        if (amount > MAX_AMOUNT) {
                            // 如果输入金额超过最大金额，则显示提示
                            amountEditText.setError("输入金额过高");
                        }
                    } catch (NumberFormatException e) {
                        // 如果无法将输入转换为数字，则清空输入框
                        amountEditText.setText("");
                    }
                }
            }
        });
    }

    // 当用户按下返回键时关闭当前 Activity，返回上一个 Activity

//    public void returnPre(View v) {
//        Intent intent = new Intent();
//        setResult(RESULT_OK, intent);
//        finish();
//    }

    @OnClick({R.id.buttonofQR})
    public void onClick(View view) {
        if (view.getId() == R.id.buttonofQR) {
           helper.goScan(this);
        }
    }

    public void initScan() {
        helper = new ActivityResultHelper(new ActivityResultHelper.OnActivityResultListener() {
            @Override
            public void onActivityResult(int requestCode, int resultCode, Intent data) {
                if (requestCode == REQUEST_CODE) {
                    // 扫描二维码/条码回传
                    if (requestCode == Constantes.REQUEST_CODE_SCAN && resultCode == RESULT_OK) {
                        if (data != null) {
                            //返回的文本内容
                            String content = data.getStringExtra(Constantes.CODED_CONTENT);
                            Bitmap bitmap = data.getParcelableExtra(Constantes.CODED_BITMAP);

                            Log.e("扫描到的内容是", "扫描到的内容是：" + content);
                            if (!TextUtils.isEmpty(content)) {
                                LogUntil.d("扫描结果： " + content);
                            }
                            if (bitmap != null) {
                            } else {
                                Log.e("扫描到的内容是", "扫描到的内容是：bitmap = null");
                            }
                        }
                    }
                }
            }
        });
    }

    public void Confirm_transfer(View v) {
        openActivity(MainActivity.class);
    }
    @Override
    public void onActivityResult(int requestCode, int resultCode, @Nullable Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        helper.onActivityResult(requestCode, resultCode, data);
    }

}
