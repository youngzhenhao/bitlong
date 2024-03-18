package com.btc.wallect.utils;

import android.content.Context;
import android.util.AttributeSet;
import android.view.MotionEvent;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.viewpager.widget.ViewPager;

public class CustomViewPager extends ViewPager {
    public CustomViewPager(@NonNull Context context) {
        super(context);
    }

    public CustomViewPager(@NonNull Context context, @Nullable AttributeSet attrs) {
        super(context, attrs);
    }

    //事件拦截
    @Override
    public boolean onInterceptTouchEvent(MotionEvent ev) {
        final int action = ev.getAction() & MotionEvent.ACTION_MASK;
        //当用户按下屏幕的那一瞬间产生该事件
        if (action == MotionEvent.ACTION_DOWN) {
            super.onInterceptTouchEvent(ev);
            //返回false表示不做拦截，事件将向下分发到子View的dispatchTouchEvent方法
            //这里就是CustomRecyclerView中重写的dispatchTouchEvent()方法
            return false;
        }
        //另外两个事件 手在屏幕上移动和抬起，
        // 事件将不再向下分发而是调用View本身的onTouchEvent方法
        return true;
    }

}
