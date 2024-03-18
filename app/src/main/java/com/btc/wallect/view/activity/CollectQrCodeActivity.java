package com.btc.wallect.view.activity;

import android.content.Intent;
import android.graphics.Bitmap;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;
import android.widget.TextView;

import com.btc.wallect.R;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.utils.CopyUtil;
import com.btc.wallect.utils.QRCodeUtil;
import com.btc.wallect.view.activity.base.BaseActivity;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class CollectQrCodeActivity extends BaseActivity {
    @BindView(R.id.img_qr_code)
    ImageView mImgQRCode;
    @BindView(R.id.tv_loction)
    TextView mTvLoction;

    String qrTxt = "这是一个测试二维码";
    private final List<CollectBean> fruitList = new ArrayList<>();

    @Override
    protected int setContentView() {
        return R.layout.act_collect_qrcode;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle("收款");
        setImgBack(true);
        setTitleBackgroundColor(3);
        setTitleTxtColor(0);
        initView();


    }


    private void initView() {

        Bitmap qrCode = QRCodeUtil.createQRCodeBitmap(qrTxt, 200, 200);
        if (qrCode != null && !qrCode.isRecycled()) {
            mImgQRCode.setImageBitmap(qrCode);
        }
    }

    @OnClick({R.id.img_new_loction, R.id.img_share, R.id.img_copy})
    public void onClick(View view) {
        if (view.getId() == R.id.img_new_loction) {

        } else if (view.getId() == R.id.img_share) {
            // 创建分享的Intent
            Intent intent = new Intent(Intent.ACTION_SEND);
            intent.setType("text/plain");
            intent.putExtra(Intent.EXTRA_TEXT, qrTxt);

            // 启动分享的活动
            startActivity(Intent.createChooser(intent, "分享到"));

        } else if (view.getId() == R.id.img_copy) {
            CopyUtil.copyClicks(qrTxt);
        }
    }
}
