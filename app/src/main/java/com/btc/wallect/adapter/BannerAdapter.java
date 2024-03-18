package com.btc.wallect.adapter;

import android.content.Context;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.viewpager.widget.PagerAdapter;

import java.util.List;



public class BannerAdapter extends PagerAdapter {

    private List<View> mViewList;
    private LayoutInflater mInflater;

    public BannerAdapter(Context context, List<View> viewList) {
        mViewList = viewList;
        mInflater = LayoutInflater.from(context);
    }

    @Override
    public int getCount() {
        return mViewList.size();
    }

    @Override
    public Object instantiateItem(ViewGroup container, int position) {
        container.addView(mViewList.get(position));
        return mViewList.get(position);
    }

    @Override
    public void destroyItem(ViewGroup container, int position, Object object) {
        container.removeView((View) object);
    }

    @Override
    public boolean isViewFromObject(View view, Object object) {
        return view == object;
    }
}
