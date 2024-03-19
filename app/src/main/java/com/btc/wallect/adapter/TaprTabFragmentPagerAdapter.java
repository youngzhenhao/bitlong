package com.btc.wallect.adapter;

import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentPagerAdapter;

import com.btc.wallect.R;
import com.btc.wallect.view.activity.fragment.TaprTabFagment;

public class TaprTabFragmentPagerAdapter extends FragmentPagerAdapter {

    private String[] mTitles = new String[]{"Tapr20", "NFT"};

    public TaprTabFragmentPagerAdapter(FragmentManager fm) {
        super(fm);
    }

    @Override
    public Fragment getItem(int position) {
        if (position == 1) {
            return new TaprTabFagment();
        } else if (position == 2) {
            return new TaprTabFagment();
        }else if (position==3){
            return new TaprTabFagment();
        }
        return new TaprTabFagment();
    }

    @Override
    public int getCount() {
        return mTitles.length;
    }

    //ViewPager与TabLayout绑定后，这里获取到PageTitle就是Tab的Text
    @Override
    public CharSequence getPageTitle(int position) {
        return mTitles[position];
    }

}
