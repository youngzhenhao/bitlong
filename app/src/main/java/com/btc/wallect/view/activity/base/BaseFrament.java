package com.btc.wallect.view.activity.base;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;

import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentTransaction;

import com.btc.wallect.db.DBOpenHelper;
import com.btc.wallect.db.DBdao;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.view.activity.MainActivity;

import java.util.List;

import butterknife.Unbinder;

public abstract class BaseFrament  extends Fragment{
    private Unbinder unbinder;
    protected boolean isVisible = false;
    private int backCount;

    private static final String STATE_SAVE_IS_HIDDEN = "STATE_SAVE_IS_HIDDEN";

    private MainActivity mActivity;
    private DBOpenHelper dBhelpUtil;
    public DBdao wallectDao;

    @Override
    public void onAttach(Context context) {
        super.onAttach(context);
        this.mActivity = (MainActivity) context;
    }

    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        SharedPreferencesHelperUtil.getInstance().init(getActivity());
        dBhelpUtil = new DBOpenHelper(getActivity(), DBOpenHelper.DB_NAME, null, DBOpenHelper.DB_VERSION);
        wallectDao = new DBdao(getActivity(), dBhelpUtil);
        backCount = mActivity.getSupportFragmentManager().getBackStackEntryCount();
//        if (backCount != 0){
//            mActivity.hideBottom();
//        }

        // 如果savedInstanceState不为空，那么可以获取在onSaveInstanceState方法中保存的值。
        if(savedInstanceState != null){
            boolean isHidden = savedInstanceState.getBoolean(STATE_SAVE_IS_HIDDEN);// 获取保存fragment前的可视状态

            FragmentTransaction transaction = getFragmentManager().beginTransaction();
            if(isHidden){
                // 如果原来fragment是隐藏状态，那么就hide
                transaction.hide(this);
            }else {
                // 如果原来fragment是显示状态，那么就show
                transaction.show(this);
            }

            transaction.commit();
        }

    }


    /**
     * 该方法在onCreate之前调用，实现数据的预加载
     * @param isVisibleToUser 界面对用户是否可见
     *                        false表示对用户不可见，反之为true
     */
    @Override
    public void setUserVisibleHint(boolean isVisibleToUser) {
        super.setUserVisibleHint(isVisibleToUser);
        if (isVisibleToUser){
            isVisible = true;
            onVisible();
        }else {
            isVisible = false;
            onInVisible();
        }
    }

    /**
     * 界面对用户不可见时，执行该方法
     */
    protected void onInVisible() {

    }

    /**
     * 界面对用户可见时，执行加载数据的方法
     */
    protected void onVisible() {
        setLazyLoad();
    }

    /**
     * 加载数据方法
     */
    protected void setLazyLoad() {

    }

    /**
     * 调用该方法用来保存fragment当前的动态状态。
     * 可以用bundle对象保存一些值，然后可以在onCreate方法中获取保存的值。
     * @param outState Bundle对象
     */
    @Override
    public void onSaveInstanceState(Bundle outState) {
        super.onSaveInstanceState(outState);
        outState.putBoolean(STATE_SAVE_IS_HIDDEN,isHidden());// 保存当前fragment的可视状态
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        backCount = mActivity.getSupportFragmentManager().getBackStackEntryCount();
//        if (backCount == 0){
//            mActivity.visibleBottom();
        }
//    }

    protected void openActivity(Class<?> Action) {
        startActivity(new Intent(getActivity(), Action));
    }
    protected void openActivityData(Class<?> Action,String toWhere) {
        Intent intent=new Intent(getActivity(),Action);
        intent.putExtra(ConStantUtil.KEY_TOACTION,toWhere);

        startActivity(intent);
    }
    protected void openActivityData(Class<?> Action, String toWhere,boolean state) {
        Intent intent = new Intent(getActivity(), Action);
        intent.putExtra(ConStantUtil.KEY_TOACTION, toWhere);
        intent.putExtra(ConStantUtil.WALLECT_STATE, state);
        startActivity(intent);
    }

    public List<Wallet> selectWallectData() {
        List<Wallet> data = wallectDao.select(null);
        if (data.equals(null) || data.size() == 0) {
            //  textView.setText("没有查到数据！");
        } else {
            //  textView.setText(data.toString());
        }
        return data;

    }

    public List<Wallet> selectDataByID(Long id) {
        List<Wallet> data = wallectDao.select(id);
        if (data.equals(null) || data.size() == 0) {

        } else {

        }
        return data;

    }

}
