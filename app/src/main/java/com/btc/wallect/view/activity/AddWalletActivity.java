package com.btc.wallect.view.activity;
import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.Spinner;

import com.btc.wallect.R;
import com.btc.wallect.view.activity.base.BaseActivity;

import butterknife.BindView;

public class AddWalletActivity extends BaseActivity {

    @BindView(R.id.mySpinner)
    Spinner spinner;

    @Override
    protected int setContentView() {
        return R.layout.add_wallet;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        ArrayAdapter<CharSequence> adapter = ArrayAdapter.createFromResource(this,
                R.array.options_array, android.R.layout.simple_spinner_item);

        // 设置下拉列表的样式
        adapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item);

        // 将适配器应用到 Spinner
        spinner.setAdapter(adapter);
    }

    // 当用户按下返回键时关闭当前 Activity，返回上一个 Activity

    public void returnPre(View v) {
        Intent intent = new Intent();
        setResult(RESULT_OK, intent);
        finish();
    }

    public void nextActivity(View v){
        Intent intent = new Intent(this, TransferActivity.class);
        startActivity(intent);
    }
}
