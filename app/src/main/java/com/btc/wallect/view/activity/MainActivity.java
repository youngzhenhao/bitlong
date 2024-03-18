package com.btc.wallect.view.activity;

import android.os.Bundle;

import android.view.KeyEvent;
import android.widget.FrameLayout;

import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.viewpager.widget.PagerAdapter;

import com.btc.wallect.view.activity.fragment.NostrFragment;
import com.btc.wallect.view.activity.fragment.TaprFragment;
import com.btc.wallect.view.activity.fragment.WallectFragment;
import com.btc.wallect.R;
import com.google.android.material.bottomnavigation.BottomNavigationView;

import butterknife.BindView;

public class MainActivity extends AppCompatActivity {


    @BindView(R.id.bottom_navigation)
    BottomNavigationView bottomNavigationView;
    @BindView(R.id.fragment_container)
    FrameLayout fragment_container;

    private PagerAdapter adapter;
    private long exitTime;
    private FragmentManager fragmentManager;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        initView();
    }

//    @Override
//    protected int setContentView() {
//        return R.layout.activity_main;
//    }
//
//    @Override
//    protected void init(View view, Bundle savedInstanceState) {
//       setTitleHide(false);
//        initView();
//        mPresenter = new MainPresentImpl(this, this);
//    }

    private void initView() {
//        fragmentManager = getSupportFragmentManager();
//
//        bottomNavigationView.setOnNavigationItemSelectedListener(navListener);
//
//        // 默认显示首页Fragment
//        fragmentManager.beginTransaction().replace(R.id.fragment_container, new WallectFragment()).commit();


        bottomNavigationView = findViewById(R.id.bottom_navigation);
        fragmentManager = getSupportFragmentManager();

        // 默认显示第一个Fragment
        fragmentManager.beginTransaction().replace(R.id.fragment_container, new WallectFragment()).commit();

        bottomNavigationView.setOnNavigationItemSelectedListener(item -> {
            Fragment selectedFragment = null;
            switch (item.getItemId()) {
                case R.id.item_wllect:
                            selectedFragment = new WallectFragment();
                            break;
                        case R.id.item_tapr:
                            selectedFragment = new TaprFragment();
                            break;
                        case R.id.item_nostr:
                            selectedFragment = new NostrFragment();
                            break;
                    }

            fragmentManager.beginTransaction().replace(R.id.fragment_container, selectedFragment).commit();
            return true;
        });

    }

//    @Override
//    public void onClick(View view) {
////        if (view.getId() == R.id.button_get){
////            mPresenter.getSearchBooks("少年维特的烦恼",null,0,1);
////        }
//    }
//
//    @Override
//    public void setBooksUISuccess(Book book) {
//        //  text_response.setText(book.toString());
//    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
    }

//    private BottomNavigationView.OnNavigationItemSelectedListener navListener =
//            new BottomNavigationView.OnNavigationItemSelectedListener() {
//                @Override
//                public boolean onNavigationItemSelected(@NonNull MenuItem item) {
//                    Fragment selectedFragment = null;
//                    switch (item.getItemId()) {
//                        case R.id.item_wllect:
//                            selectedFragment = new WallectFragment();
//                            break;
//                        case R.id.item_tapr:
//                            selectedFragment = new TaprFragment();
//                            break;
//                        case R.id.item_nostr:
//                            selectedFragment = new NostrFragment();
//                            break;
//                    }
//                    fragmentManager.beginTransaction().replace(R.id.fragment_container, selectedFragment).commit();
//                    return true;
//                }
//            };

    @Override
    public boolean dispatchKeyEvent(KeyEvent event) {
        if (event.getKeyCode() == KeyEvent.KEYCODE_BACK
                && event.getAction() == KeyEvent.ACTION_DOWN
                && event.getRepeatCount() == 0) {
            // 重写键盘事件分发，onKeyDown方法某些情况下捕获不到，只能在这里写
            if ((System.currentTimeMillis() - exitTime) > 2000) {
                exitTime = System.currentTimeMillis();
            } else {
                finish();
            }
            return true;
        }
        return super.dispatchKeyEvent(event);
    }


}
