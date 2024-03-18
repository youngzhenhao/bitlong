package com.btc.wallect.utils;

import android.content.Context;
import android.util.AttributeSet;
import android.view.MotionEvent;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.recyclerview.widget.RecyclerView;

public class CustomRecyclerView extends RecyclerView {
    private int mLastX;
    private int mLastY;

    public CustomRecyclerView(@NonNull Context context) {
        super(context);
    }

    public CustomRecyclerView(@NonNull Context context, @Nullable AttributeSet attrs) {
        super(context, attrs);
    }

    public CustomRecyclerView(@NonNull Context context, @Nullable AttributeSet attrs, int defStyle) {
        super(context, attrs, defStyle);
    }

    //处理触摸事件的分发 是从dispatchTouchEvent开始的
    @Override
    public boolean dispatchTouchEvent(MotionEvent event) {
        //触摸点相对于其所在组件原点的X坐标
        int x = (int) event.getX();
        int y = (int) event.getY();
        switch (event.getAction()) {
            case MotionEvent.ACTION_DOWN:
                //手按下屏幕,父布局没有作用,进行拦截
                //让父布局ViewPager禁用拦截功能,从而让父布局忽略事件后的一切行为
                //requestDisallowInterceptTouchEvent(true)表示：
                //getParent() 获取到父视图 父视图不拦截触摸事件
                //孩子不希望父视图拦截触摸事件
                getParent().requestDisallowInterceptTouchEvent(true);
                break;
            case MotionEvent.ACTION_MOVE:
                //水平移动的增量
                int deltaX = x - mLastX;
                int deltaY = y - mLastY;
                //Math.abs绝对值
                if (Math.abs(deltaX) > Math.abs(deltaY)) {
                    //当水平增量大于竖直增量时，表示水平滑动，此时需要父View去处理事件，所以不拦截
                    //让父布局ViewPager使用拦截功能,从而让父布局完成事件后的一切行为

                    //requestDisallowInterceptTouchEvent(false)表示：
                    //孩子希望父视图拦截触摸事件,也就是让CustomViewPager拦截触摸事件，进行左右滑动
                    getParent().requestDisallowInterceptTouchEvent(false);
                }
                break;
            default:
                break;
        }
        mLastX = x;
        mLastY = y;
        return super.dispatchTouchEvent(event);
    }

}
