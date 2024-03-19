package com.btc.wallect.view.activity;

import android.content.Intent;
import android.os.Bundle;
import android.text.Editable;
import android.text.TextWatcher;
import android.view.View;
import android.view.WindowManager;
import android.widget.EditText;


import androidx.appcompat.app.AppCompatActivity;

import com.btc.wallect.R;
import com.btc.wallect.view.activity.base.BaseActivity;

import butterknife.BindView;

public class TransferActivity extends BaseActivity {

    private static final int MAX_AMOUNT = 1000; // 设置允许的最大金额
@BindView( R.id.ed_text_amount)
EditText amountEditText;
    @Override
    protected int setContentView() {
        return R.layout.transfer;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {


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

    public void returnPre(View v) {
        Intent intent = new Intent();
        setResult(RESULT_OK, intent);
        finish();
    }

    public void startScan(View v) {
//       //设置竖屏
//       setRequestedOrientation(ActivityInfo.SCREEN_ORIENTATION_PORTRAIT);
//       // 创建 IntentIntegrator 对象
//       IntentIntegrator integrator = new IntentIntegrator(this);
//       // 设置自定义提示信息
//       integrator.setPrompt("请将二维码放入框内");
//       // 启动扫描
//       integrator.initiateScan();
    }

    public void Confirm_transfer(View v) {
        openActivity(MainActivity.class);
    }

}
