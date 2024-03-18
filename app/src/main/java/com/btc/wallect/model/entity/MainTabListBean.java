package com.btc.wallect.model.entity;

import java.util.List;

public class MainTabListBean {
    private String tabTxt;
    private boolean isSelect;
    private List<MainDateilBean> dataTabList;

    public String getTabTxt() {
        return tabTxt;
    }

    public void setTabTxt(String tabTxt) {
        this.tabTxt = tabTxt;
    }

    public boolean isSelect() {
        return isSelect;
    }

    public void setSelect(boolean select) {
        isSelect = select;
    }

    public List<MainDateilBean> getDataTabList() {
        return dataTabList;
    }

    public void setDataTabList(List<MainDateilBean> dataTabList) {
        this.dataTabList = dataTabList;
    }
}
