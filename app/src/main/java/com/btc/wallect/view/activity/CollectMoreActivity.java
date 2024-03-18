package com.btc.wallect.view.activity;

import android.animation.ObjectAnimator;
import android.content.Intent;
import android.graphics.Bitmap;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RadioGroup;
import android.widget.TextView;

import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.adapter.HistoryAdapter;
import com.btc.wallect.model.entity.HistoryBean;
import com.btc.wallect.utils.CopyUtil;
import com.btc.wallect.utils.DateTimeUtil;
import com.btc.wallect.utils.QRCodeUtil;
import com.btc.wallect.utils.ScreenHelper;
import com.btc.wallect.utils.ToastUtils;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.R;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;

public class CollectMoreActivity extends BaseActivity {

//    @BindView(R.id.app_bar)
//    AppBarLayout mAppBar;

    @BindView(R.id.search_tab_container)
    View slideView;

    //
    @BindView(R.id.rg_slide)
    RadioGroup rgSlide;
    @BindView(R.id.slide_bg)
    View mSlideView;
    @BindView(R.id.recycler_history_detail)
    RecyclerView mRecycler_history_detail;
    //    @BindView(R.id.ll_tab_left_1)
//    LinearLayout mLl_tab_left_1;
    @BindView(R.id.ll_tab_left_2)
    LinearLayout mLl_tab_left_2;
    //    @BindView(R.id.ll_tab_center_1)
//    LinearLayout mLl_tab_center_1;
//    @BindView(R.id.ll_tab_right_1)
    LinearLayout mLl_tab_right_1;
    @BindView(R.id.img_qr_code)
    ImageView mImgQRCode;
    @BindView(R.id.tv_loction)
    TextView mTvLoction;
    @BindView(R.id.top_container_1)
    LinearLayout mTop_container_1;
    @BindView(R.id.top_container_2)
    LinearLayout mTop_container_2;
    @BindView(R.id.top_container_3)
    LinearLayout mTop_container_3;
    @BindView(R.id.top_container_4)
    LinearLayout mTop_container_4;
    @BindView(R.id.tv_submit)
    TextView mTvSubmit;
    @BindView(R.id.img_right_qr_code)
            ImageView mImgRightQRCode;

    String qrTxt = "这是一个测试二维码";
    private int underLineWidth = 0;
    public List<HistoryBean> historyBeanList;

    @Override
    protected int setContentView() {
        return R.layout.act_collect_more;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setImgBack(true);
        setTitle("收款");
        initView();
        slide(0);
        initCollectTest();
        setRecyclerMianDatailView();
    }

    private void initView() {
        historyBeanList = new ArrayList<>();
        rgSlide.check(R.id.rb_left);
//        mAppBar.addOnOffsetChangedListener(new AppBarLayout.OnOffsetChangedListener() {
//            @Override
//            public void onOffsetChanged(AppBarLayout appBarLayout, int verticalOffset) {
//                slideView.setAlpha(getAlpha(verticalOffset, appBarLayout.getTotalScrollRange()));
//                mTop.setAlpha(1.0f - getAlpha(verticalOffset, appBarLayout.getTotalScrollRange()));
//            }
//        });
        rgSlide.setOnCheckedChangeListener(new RadioGroup.OnCheckedChangeListener() {
            @Override
            public void onCheckedChanged(RadioGroup group, int checkedId) {
                switch (checkedId) {
                    case R.id.rb_left:
                        mSlideView.setBackgroundResource(R.drawable.bg_tab_left);
                        slide(0);
                        setTAB(0);
                        break;
                    case R.id.rb_center:
                        mSlideView.setBackgroundResource(R.drawable.bg_tab_center);
                        slide(1);
                        setTAB(1);
                        break;
                    case R.id.rb_right:
                        slide(2);
                        setTAB(2);
                        mSlideView.setBackgroundResource(R.drawable.bg_tab_right);
                        break;
                }
            }
        });
        Bitmap qrCode = QRCodeUtil.createQRCodeBitmap(qrTxt, 200, 200);
        if (qrCode != null && !qrCode.isRecycled()) {
            mImgQRCode.setImageBitmap(qrCode);
        }
    }

    private float getAlpha(int a, int b) {
        float f = Math.abs((float) a / b);
        float result = 1.0f - f;
        return result;
    }


    private void slide(int position) {
        if (underLineWidth <= 0) {
            underLineWidth = (ScreenHelper.getScreenWidth() - ScreenHelper.dip2px(20)) / 3;
            LinearLayout.LayoutParams layoutParams = (LinearLayout.LayoutParams) mSlideView.getLayoutParams();
            layoutParams.width = underLineWidth;
            mSlideView.setLayoutParams(layoutParams);
        }
        ObjectAnimator.ofFloat(mSlideView, "translationX", ScreenHelper.dip2px(10) + underLineWidth * position).setDuration(300).start();
    }

    private void setRecyclerMianDatailView() {

        LinearLayoutManager layoutManager = new LinearLayoutManager(this);
        mRecycler_history_detail.setLayoutManager(layoutManager);
        View headerView = getLayoutInflater().inflate(R.layout.item_history_head, mRecycler_history_detail, false);

        HistoryAdapter historyAdapter = new HistoryAdapter(historyBeanList, headerView);
        mRecycler_history_detail.setAdapter(historyAdapter);

    }

    private void initCollectTest() {

        for (int i = 0; i < 10; i++) {
            HistoryBean historyBean = new HistoryBean();
            historyBean.setAmount(0.1 + i);
            historyBean.setState(0);
            historyBean.setDataTime(DateTimeUtil.getInstance().getNowTime());
            historyBeanList.add(historyBean);
        }


    }

    private void setTAB(int index) {
        switch (index) {
            case 0:

                mLl_tab_left_2.setVisibility(View.VISIBLE);

                mTop_container_1.setVisibility(View.VISIBLE);
                mTop_container_2.setVisibility(View.GONE);
                mTop_container_3.setVisibility(View.GONE);
                mTop_container_4.setVisibility(View.GONE);
                setTitleBackgroundColor(1);
                setTitleTxtColor(2);
                break;
            case 1:
                mLl_tab_left_2.setVisibility(View.GONE);
                mTop_container_1.setVisibility(View.GONE);
                mTop_container_2.setVisibility(View.VISIBLE);
                mTop_container_3.setVisibility(View.GONE);
                mTop_container_4.setVisibility(View.GONE);
                setTitleBackgroundColor(4);
                setTitleTxtColor(2);
                break;
            case 2:

                mLl_tab_left_2.setVisibility(View.GONE);
                mTop_container_1.setVisibility(View.GONE);
                mTop_container_2.setVisibility(View.GONE);
                mTop_container_3.setVisibility(View.VISIBLE);
                mTop_container_4.setVisibility(View.GONE);
                setTitleBackgroundColor(1);
                break;
            case 3:
                mLl_tab_left_2.setVisibility(View.GONE);
                mTop_container_1.setVisibility(View.GONE);
                mTop_container_2.setVisibility(View.GONE);
                mTop_container_3.setVisibility(View.GONE);
                mTop_container_4.setVisibility(View.VISIBLE);
                Bitmap qrCode = QRCodeUtil.createQRCodeBitmap(qrTxt, 200, 200);
                if (qrCode != null && !qrCode.isRecycled()) {
                    mImgRightQRCode.setImageBitmap(qrCode);
                }
                break;
        }
    }


    @OnClick({R.id.img_share, R.id.img_copy, R.id.tv_submit,R.id.img_right_share,R.id.img_right_copy})
    public void onClick(View view) {
        if (view.getId() == R.id.img_share) {
            setShare();

        } else if (view.getId() == R.id.img_copy) {
            CopyUtil.copyClicks(qrTxt);
        } else if (view.getId() == R.id.tv_submit) {
            runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    ToastUtils.initToast("地址已存在", 1);
                }
            });

            setTAB(3);
        }
        else if (view.getId() == R.id.img_right_share) {
            setShare();
        }
        else if (view.getId() == R.id.img_right_copy) {
            CopyUtil.copyClicks(qrTxt);
        }
    }

    private void setShare(){
        // 创建分享的Intent
        Intent intent = new Intent(Intent.ACTION_SEND);
        intent.setType("text/plain");
        intent.putExtra(Intent.EXTRA_TEXT, qrTxt);
        // 启动分享的活动
        startActivity(Intent.createChooser(intent, "分享到"));
    }
}
