package com.btc.wallect.view.activity;

import android.content.ClipData;
import android.content.ClipDescription;
import android.content.ClipboardManager;
import android.content.Context;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;

import com.btc.wallect.R;
import com.btc.wallect.utils.DialogUtil;
import com.btc.wallect.view.activity.base.BaseActivity;

import butterknife.BindView;
import butterknife.OnClick;

public class ImportKeyAcivity extends BaseActivity {
    @BindView(R.id.ed_key_txt)
    EditText mEdKeyTxt;

    @Override
    protected int setContentView() {
        return R.layout.act_import_key;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle(" ");
        setImgBack(true);
        setTitleBackgroundColor(2);
    }

    @OnClick({R.id.tv_submit, R.id.tv_paste})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_submit) {
         //   DialogUtil.showSimpleDialog(this, "提示", "确认", null);
          openActivity(AddWalletActivity.class);
        } else if (view.getId() == R.id.tv_paste) {
            ClipboardManager clipboard = (ClipboardManager) getSystemService(Context.CLIPBOARD_SERVICE);
            if (clipboard.hasPrimaryClip() && clipboard.getPrimaryClipDescription().hasMimeType(ClipDescription.MIMETYPE_TEXT_PLAIN)) {

                ClipData.Item item = clipboard.getPrimaryClip().getItemAt(0);
                CharSequence text = item.getText();
                mEdKeyTxt.setText(text);
            } else {
                // 粘贴板没有文本内容
            }
        }
    }
}
