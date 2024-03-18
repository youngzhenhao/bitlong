package com.btc.wallect.adapter;



import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentPagerAdapter;

import com.btc.wallect.view.activity.fragment.NostrFragment;
import com.btc.wallect.view.activity.fragment.TaprFragment;
import com.btc.wallect.view.activity.fragment.WallectFragment;


/**
 * 主界面底部菜单适配器
 */
public class MainFragmentAdapter extends FragmentPagerAdapter {
    public MainFragmentAdapter(FragmentManager fm) {
        super(fm);
    }

    @Override
    public Fragment getItem(int i) {
        Fragment fragment = null;
        switch (i) {
            case 0:
                fragment = new WallectFragment();
                break;
            case 1:
                fragment = new TaprFragment();
                break;
            case 2:
                fragment = new NostrFragment();
                break;

            default:
                break;
        }
        return fragment;
    }

    @Override
    public int getCount() {
        return 3;
    }

}
