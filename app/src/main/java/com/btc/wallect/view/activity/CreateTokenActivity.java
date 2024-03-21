package com.btc.wallect.view.activity;

import android.content.Intent;
import android.graphics.Color;
import android.net.Uri;
import android.os.Bundle;
import android.provider.MediaStore;
import android.text.Spannable;
import android.text.SpannableString;
import android.text.style.ForegroundColorSpan;
import android.view.View;
import android.view.WindowManager;
import android.widget.TextView;

import androidx.appcompat.app.AppCompatActivity;

import com.btc.wallect.R;
import com.btc.wallect.view.activity.base.BaseActivity;

public class CreateTokenActivity extends BaseActivity {
    private int PICK_IMAGE_REQUEST_CODE = 1;

    @Override
    protected int setContentView() {
        return R.layout.create_token;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle("创建");
        setTitleTxtColor(2);
        setTitleBackgroundColor(1);
        setImgBack(true);
//        getWindow().setFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN,
//                WindowManager.LayoutParams.FLAG_FULLSCREEN);
//        setContentView(R.layout.create_token);
        TextView textView1 = findViewById(R.id.logo);
        TextView textView2 = findViewById(R.id.name);
        TextView textView3 = findViewById(R.id.type);
        TextView textView4 = findViewById(R.id.add);
        setRequiredText(textView4,"*是否可增发");
        TextView textView5 = findViewById(R.id.introduce);
        TextView textView6 = findViewById(R.id.quantity);
        setRequiredText(textView6,"*数量");
        setRequiredText(textView5,"*介绍");
        setRequiredText(textView1,"*上传logo");
        setRequiredText(textView2,"*名称");
        setRequiredText(textView3,"*类别");
    }

    private void setRequiredText(TextView textView, String text) {
        SpannableString spannableString = new SpannableString(text);
        int starIndex = text.indexOf("*");
        if (starIndex != -1) {
            spannableString.setSpan(new ForegroundColorSpan(Color.parseColor("#EC3468")), starIndex, starIndex + 1, Spannable.SPAN_EXCLUSIVE_EXCLUSIVE);
        }
        textView.setText(spannableString);
    }



    public void returnPre(View v) {
        Intent intent = new Intent();
        setResult(RESULT_OK, intent);
        finish();
    }

    public void openGallery(View view) {
        Intent intent = new Intent(Intent.ACTION_PICK, MediaStore.Images.Media.EXTERNAL_CONTENT_URI);
        String[] mimeTypes = {"image/png", "image/jpeg"}; // 定义要支持的 MIME 类型
        intent.putExtra(Intent.EXTRA_MIME_TYPES, mimeTypes); // 将 MIME 类型数组传递给 Intent
        startActivityForResult(intent, PICK_IMAGE_REQUEST_CODE);
    }
    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);

        if (requestCode == PICK_IMAGE_REQUEST_CODE && resultCode == RESULT_OK && data != null && data.getData() != null) {
            Uri selectedImageUri = data.getData();

            // 在这里处理选择的图片

        }
    }


}
