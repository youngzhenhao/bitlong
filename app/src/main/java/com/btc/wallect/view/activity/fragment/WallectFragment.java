package com.btc.wallect.view.activity.fragment;

import android.graphics.Color;
import android.os.Bundle;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.TextView;


import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.adapter.AlbumPanoramaAdapter;
import com.btc.wallect.adapter.MainBtcAdapter;
import com.btc.wallect.adapter.OurAdapter;
import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.model.entity.MainDateilBean;
import com.btc.wallect.model.entity.MainTabListBean;
import com.btc.wallect.model.entity.WalletListBean;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.GsonUtils;
import com.btc.wallect.utils.LogUntil;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.utils.dialog.ChannelOpeningDialog;
import com.btc.wallect.utils.dialog.WallectDialog;
import com.btc.wallect.view.activity.TokenDetailActivity;
import com.btc.wallect.view.activity.WallectEditActivity;
import com.btc.wallect.view.activity.base.BaseFrament;
import com.google.gson.Gson;
import com.jude.rollviewpager.OnItemClickListener;
import com.jude.rollviewpager.RollPagerView;
import com.jude.rollviewpager.hintview.ColorPointHintView;


import java.util.ArrayList;
import java.util.List;


public class WallectFragment extends BaseFrament implements View.OnClickListener {
//    @BindView(R.id.recycler_main_view)
//    RecyclerView mRerecyclerMainview;
//    @BindView(R.id.img_btc_hide)
//    ImageView mImg_btc_hide;


    private static final int REQUEST_EXTERNAL_STORAGE = 1;
    private MainBtcAdapter mainBtcAdapter;
    private AlbumPanoramaAdapter mainBtcTabAdapter;
    public List<MainDateilBean> dataList;
    public List<MainTabListBean> mMainTabList;
    private RecyclerView.LayoutManager layoutManager;

    private RecyclerView mRerecyclerMainview;
    private RecyclerView mRecycler_main_tabl;
    private LinearLayout mLl_main_tab;
    private RollPagerView mBanner;
    private List imagesList = new ArrayList();
    private ImageView mImg_btc_hide, mImgWallect;
    private TextView mTv_btc_datail;
    private boolean isShowNum = true;
    public List<WalletListBean> walletListBeans;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_wallect, container, false);

        initViews(view);
        return view;
    }


    private void initViews(View view) {
        mRerecyclerMainview = view.findViewById(R.id.recycler_main_view);
        mRecycler_main_tabl = view.findViewById(R.id.recycler_main_tab);
        mLl_main_tab = view.findViewById(R.id.ll_main_tab);
        mBanner = view.findViewById(R.id.banner);
        mImg_btc_hide = view.findViewById(R.id.img_btc_hide);
        mImgWallect = view.findViewById(R.id.img_wallect);
        mTv_btc_datail=view.findViewById(R.id.tv_btc_datail);
        mLl_main_tab.setOnClickListener(this);
        mImg_btc_hide.setOnClickListener(this);
        mImgWallect.setOnClickListener(this);
        mTv_btc_datail.setOnClickListener(this);
        dataList = new ArrayList<>();
        mMainTabList = new ArrayList<>();
        walletListBeans = new ArrayList<>();
        initCollectTest();
        setTab();
        setRecyclerMianDatailTabView();
        setRecyclerMianDatailView();
        initBanner();
        isShowDatail();
        setWallectDatail();
    }

    private void setRecyclerMianDatailTabView() {
        String tabList = SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.MAIN_TAB_LIST, "");
        LogUntil.d(tabList);
        mMainTabList = GsonUtils.jsonToList(tabList, MainTabListBean.class);

        mRecycler_main_tabl.setLayoutManager(new LinearLayoutManager(getActivity(), LinearLayoutManager.HORIZONTAL, false));
        mainBtcTabAdapter = new AlbumPanoramaAdapter(mMainTabList, getActivity());
        mainBtcTabAdapter.addFooterView(LayoutInflater.from(getActivity()).inflate(R.layout.item_main_btc_tab_newbtn, null));

        mRecycler_main_tabl.setAdapter(mainBtcTabAdapter);
        mainBtcTabAdapter.setonItemClickListener(new onItemClickListener() {
            @Override
            public void onItemClick(int position, String txt) {

                String tabList = SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.MAIN_TAB_LIST, "");
                mMainTabList = GsonUtils.jsonToList(tabList, MainTabListBean.class);
                dataList = mMainTabList.get(position).getDataTabList();
                setRecyclerMianDatailView();


            }

            @Override
            public void onItemClick() {
                ChannelOpeningDialog openingDialog = new ChannelOpeningDialog(getActivity());
                openingDialog.setOnSureClickListenerr(new ChannelOpeningDialog.onSureClickListener() {
                    @Override
                    public void onOkClick() {
                        String tabList = SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.MAIN_TAB_LIST, "");
                        mMainTabList = GsonUtils.jsonToList(tabList, MainTabListBean.class);
                        MainTabListBean mainTabListBean = new MainTabListBean();
                        mainTabListBean.setTabTxt("通道" + mMainTabList.size());
                        mainTabListBean.setSelect(false);
                        MainDateilBean collectBean6 = new MainDateilBean();
                        collectBean6.setBtcName("BTC111");
                        collectBean6.setBtcMode("BRC20222");
                        collectBean6.setBtcAmount(0.1);
                        collectBean6.setBtcAll(22);
                        List<MainDateilBean> dataList = new ArrayList<>();
                        dataList.add(collectBean6);
                        mainTabListBean.setDataTabList(dataList);
                        mMainTabList.add(mainTabListBean);
                        SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.MAIN_TAB_LIST, new Gson().toJson(mMainTabList));
                        setRecyclerMianDatailTabView();
                    }
                });
                openingDialog.show();


            }
        });

    }

    private void setRecyclerMianDatailView() {

        layoutManager = new LinearLayoutManager(getActivity());
        mRerecyclerMainview.setLayoutManager(layoutManager);
        mainBtcAdapter = new MainBtcAdapter(dataList);
        mRerecyclerMainview.setAdapter(mainBtcAdapter);
        mainBtcAdapter.setonItemClickListener(new onItemClickListener() {
            @Override
            public void onItemClick(int position, String txt) {

            }

            @Override
            public void onItemClick() {
                openActivity(TokenDetailActivity.class);
            }
        });
    }


    private void initCollectTest() {
        dataList.clear();
        for (int i = 0; i < 10; i++) {
            if (i == 0) {
                MainDateilBean collectBean1 = new MainDateilBean();
                collectBean1.setBtcName("BTC");
                collectBean1.setBtcMode("BRC20");
                collectBean1.setBtcAmount(0.5 + i);
                collectBean1.setBtcAll(i);
                dataList.add(collectBean1);
            } else if (i == 1) {
                MainDateilBean collectBean2 = new MainDateilBean();
                collectBean2.setBtcName("BTC");
                collectBean2.setBtcMode("BRC20");
                collectBean2.setBtcAmount(0.6 + i);
                collectBean2.setBtcAll(i);
                dataList.add(collectBean2);
            } else if (i == 2) {
                MainDateilBean collectBean3 = new MainDateilBean();
                collectBean3.setBtcName("BTC");
                collectBean3.setBtcMode("BRC20");
                collectBean3.setBtcAmount(0.7 + i);
                collectBean3.setBtcAll(i);
                dataList.add(collectBean3);
            } else if (i == 3) {
                MainDateilBean collectBean4 = new MainDateilBean();
                collectBean4.setBtcName("BTC");
                collectBean4.setBtcMode("BRC20");
                collectBean4.setBtcAmount(0.8 + i);
                dataList.add(collectBean4);
            } else if (i == 4) {
                MainDateilBean collectBean5 = new MainDateilBean();
                collectBean5.setBtcName("BTC");
                collectBean5.setBtcMode("BRC20");
                collectBean5.setBtcAmount(0.9 + i);
                dataList.add(collectBean5);
            } else if (i == 5) {
                MainDateilBean collectBean6 = new MainDateilBean();
                collectBean6.setBtcName("BTC");
                collectBean6.setBtcMode("BRC20");
                collectBean6.setBtcAmount(0.1 + i);
                collectBean6.setBtcAll(i);
                dataList.add(collectBean6);
            } else {
                MainDateilBean collectBean6 = new MainDateilBean();
                collectBean6.setBtcName("BTC111");
                collectBean6.setBtcMode("BRC20222");
                collectBean6.setBtcAmount(0.1 + i);
                collectBean6.setBtcAll(i);
                dataList.add(collectBean6);
            }

        }


    }

    private void setTab() {

        MainTabListBean mainTabListBean = new MainTabListBean();
        mainTabListBean.setTabTxt("通道0");
        mainTabListBean.setSelect(false);

        List<MainDateilBean> dataList = new ArrayList<>();
        MainDateilBean collectBean6 = new MainDateilBean();
        collectBean6.setBtcName("BTC111");
        collectBean6.setBtcMode("BRC20222");
        collectBean6.setBtcAmount(0.1);
        collectBean6.setBtcAll(22);
        dataList.add(collectBean6);

        mainTabListBean.setDataTabList(dataList);


        mMainTabList.add(mainTabListBean);
        LogUntil.d(new Gson().toJson(mMainTabList));

        SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.MAIN_TAB_LIST, GsonUtils.listTojson2(mMainTabList));
    }


    @Override
    public void onClick(View v) {
        if (v.getId() == R.id.ll_main_tab) {
            initCollectTest();
            setRecyclerMianDatailView();
        } else if (v.getId() == R.id.img_btc_hide) {
            isShowDatail();
        } else if (v.getId() == R.id.img_wallect) {
            setDialogList();
        } else if (v.getId()==R.id.tv_btc_datail) {
            openActivity(WallectEditActivity.class);
        }
    }

    private void initBanner() {
        imagesList.add(R.mipmap.img_banner);
        imagesList.add(R.mipmap.img_crate_wallet_one);

        //设置轮播图下的点的颜色
        mBanner.setHintView(new ColorPointHintView(getActivity(), Color.RED, Color.WHITE));
        //设置图片轮播时间
        mBanner.setPlayDelay(2000);
        //设置动画的持续时间
        //设置动画后很别扭，取消与否在与自己
        // mCarousel.setAnimationDurtion(2000);
        mBanner.setAnimationDurtion(0);
        //设置适配器
        mBanner.setAdapter(new OurAdapter(imagesList));
        //子条目点击
        mBanner.setOnItemClickListener(new OnItemClickListener() {
            @Override
            public void onItemClick(int position) {


            }
        });
    }

    private void isShowDatail() {
        if (isShowNum) {
            mImg_btc_hide.setImageResource(R.mipmap.icon_btc_hide1);
            isShowNum = false;
        } else {
            mImg_btc_hide.setImageResource(R.mipmap.icon_btc_hide2);
            isShowNum = true;
        }
    }

    private void setWallectDatail() {
        for (int i = 0; i < 5; i++) {
            WalletListBean walletListBean = new WalletListBean();
            walletListBean.setWallectAmount(0.222);
            walletListBean.setWallectName("BTC-1" + i);
            walletListBean.setWallectKey("bc362....2dfsvd1");
            walletListBeans.add(walletListBean);
        }

    }

    private void setDialogList() {
        WallectDialog wallectDialog = new WallectDialog(getActivity(), walletListBeans);
        wallectDialog.setAddOnclickListener(new WallectDialog.onAddClickListener() {
            @Override
            public void onAddClick() {

            }
        });
        wallectDialog.show();
    }

}