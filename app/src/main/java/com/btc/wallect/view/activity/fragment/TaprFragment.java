package com.btc.wallect.view.activity.fragment;

import android.graphics.Color;
import android.graphics.Typeface;
import android.os.Bundle;

import android.util.TypedValue;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;


import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import androidx.viewpager.widget.ViewPager;

import com.btc.wallect.R;
import com.btc.wallect.adapter.HistoryAdapter;
import com.btc.wallect.adapter.MsgContentFragmentAdapter;
import com.btc.wallect.adapter.OurAdapter;
import com.btc.wallect.adapter.TaprTabAdapter;
import com.btc.wallect.adapter.TaprTabFragmentPagerAdapter;
import com.btc.wallect.model.entity.HistoryBean;
import com.btc.wallect.model.entity.TaprTabListBean;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.DateTimeUtil;
import com.btc.wallect.view.activity.CreateWalletActivity;
import com.btc.wallect.view.activity.ImportKeyAcivity;
import com.google.android.material.tabs.TabLayout;
import com.jude.rollviewpager.OnItemClickListener;
import com.jude.rollviewpager.RollPagerView;
import com.jude.rollviewpager.hintview.ColorPointHintView;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;
import butterknife.OnClick;


public class TaprFragment extends Fragment {
    @BindView(R.id.tapr_tabLayout)
    TabLayout mTabLayout;
    @BindView(R.id.tapr_viewPager)
    ViewPager mViewPager;
    @BindView(R.id.roll_banner)
    RollPagerView mRoll_banner;
    @BindView(R.id.tv_me_attention_list)
    TextView mTv_me_attention_list;
    @BindView(R.id.tv_hot_list)
    TextView mTv_hot_list;
    @BindView(R.id.tv_all_list)
    TextView mTv_all_list;
    @BindView(R.id.recycler_tapr_detail)
    RecyclerView mRecyclerTaprDetail;


    private MsgContentFragmentAdapter adapter;
    private List<String> names;
    private List imagesList = new ArrayList();
    private TaprTabFragmentPagerAdapter myFragmentPagerAdapter;

    private TabLayout.Tab one;
    private TabLayout.Tab two;
private TaprTabAdapter taprTabAdapter;
    public List<TaprTabListBean> taprTabList;
    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);


    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_tapr, container, false);
        ButterKnife.bind(this, view);
        taprTabList = new ArrayList<>();

        initBanner();
        //初始化视图
        initViews();
        setStatelist(1);
        initTabTest();
        setTaprTabAdapter();
        return view;
    }


    private void initBanner() {
        imagesList.add(R.mipmap.img_banner);
        imagesList.add(R.mipmap.img_banner2);

        //设置轮播图下的点的颜色
        mRoll_banner.setHintView(new ColorPointHintView(getActivity(), Color.RED, Color.WHITE));
        //设置图片轮播时间
        mRoll_banner.setPlayDelay(2000);
        //设置动画的持续时间
        //设置动画后很别扭，取消与否在与自己
        // mCarousel.setAnimationDurtion(2000);
        mRoll_banner.setAnimationDurtion(0);
        //设置适配器
        mRoll_banner.setAdapter(new OurAdapter(imagesList));
        //子条目点击
        mRoll_banner.setOnItemClickListener(new OnItemClickListener() {
            @Override
            public void onItemClick(int position) {


            }
        });
    }

    private void initViews() {


        myFragmentPagerAdapter = new TaprTabFragmentPagerAdapter(getChildFragmentManager());


        mViewPager.setAdapter(myFragmentPagerAdapter);

        //将TabLayout与ViewPager绑定在一起
        mTabLayout.setupWithViewPager(mViewPager);
        // changeTabsFont();
        //指定Tab的位置
        one = mTabLayout.getTabAt(0);
        two = mTabLayout.getTabAt(1);

    }
    private void setTaprTabAdapter(){


        LinearLayoutManager layoutManager = new LinearLayoutManager(getActivity());
        mRecyclerTaprDetail.setLayoutManager(layoutManager);
        View headerView = getLayoutInflater().inflate(R.layout.item_history_head, mRecyclerTaprDetail, false);
        TaprTabAdapter taprTabAdapter1 = new TaprTabAdapter(getActivity(),taprTabList, headerView);
        mRecyclerTaprDetail.setNestedScrollingEnabled(false);
        mRecyclerTaprDetail.setAdapter(taprTabAdapter1);
    }

    @OnClick({R.id.tv_me_attention_list, R.id.tv_hot_list, R.id.tv_all_list, R.id.ll_btn_create})
    public void onClick(View view) {
        if (view.getId() == R.id.tv_me_attention_list) {
            setStatelist(1);
        } else if (view.getId() == R.id.tv_hot_list) {
            setStatelist(2);
        } else if (view.getId() == R.id.tv_all_list) {
            setStatelist(3);
        } else if (view.getId() == R.id.ll_btn_create) {

        }

    }

    private void setStatelist(int pistion) {
      switch (pistion){
          case 1:
              mTv_me_attention_list.setTextColor(Color.parseColor("#383838"));
              mTv_hot_list.setTextColor(Color.parseColor("#808080"));
              mTv_all_list.setTextColor(Color.parseColor("#808080"));
              break;
          case 2:
              mTv_me_attention_list.setTextColor(Color.parseColor("#808080"));
              mTv_hot_list.setTextColor(Color.parseColor("#383838"));
              mTv_all_list.setTextColor(Color.parseColor("#808080"));
              break;
          case 3:
              mTv_me_attention_list.setTextColor(Color.parseColor("#808080"));
              mTv_hot_list.setTextColor(Color.parseColor("#808080"));
              mTv_all_list.setTextColor(Color.parseColor("#383838"));
              break;
      }
    }
    private void initTabTest() {

        for (int i = 0; i < 15; i++) {
            if(i==0){
                TaprTabListBean taprTabListBean = new TaprTabListBean();
                taprTabListBean.setDealName("stax");
                taprTabListBean.setState(0);
                taprTabListBean.setDealAmount(663.14+i);
                taprTabListBean.setDealPrice(6.15151515);
                taprTabListBean.setDealdDetailPrice(66314);
                taprTabListBean.setRose(663.222);
                taprTabList.add(taprTabListBean);
            }else {
                TaprTabListBean taprTabListBean = new TaprTabListBean();
                taprTabListBean.setDealName("stax");
                taprTabListBean.setState(1);
                taprTabListBean.setDealAmount(663.14+i);
                taprTabListBean.setDealPrice(6.15151515);
                taprTabListBean.setDealdDetailPrice(66314);
                taprTabListBean.setRose(663.222);
                taprTabList.add(taprTabListBean);
            }

        }


    }

}
