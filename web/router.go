package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"wechat-mall-backend/app/interfaces"
)

type Router struct {
	*mux.Router
	ch *alice.Chain
}

func NewRouter() *Router {
	mw := Middleware{}
	chain := alice.New(mw.RequestTraceHandler, mw.LoggingHandler, mw.CORSHandler, mw.ValidateAuthToken)
	return &Router{
		Router: mux.NewRouter(),
		ch:     &chain,
	}
}

func (r *Router) registerHandler(services *interfaces.MallHttpServiceImpl) {
	r.Handle("/api/wxapp/login", r.ch.ThenFunc(services.Login)).Methods("GET").Queries("code", "{code}")
	r.Handle("/api/wxapp/user-info", r.ch.ThenFunc(services.UserInfo)).Methods("GET")
	r.Handle("/api/wxapp/auth-phone", r.ch.ThenFunc(services.AuthPhone)).Methods("POST")
	r.Handle("/api/wxapp/auth-info", r.ch.ThenFunc(services.AuthUserInfo)).Methods("POST")
	r.Handle("/api/home/banner", r.ch.ThenFunc(services.HomeBanner)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/home/grid", r.ch.ThenFunc(services.GetGridCategoryList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/category/list", r.ch.ThenFunc(services.GetSubCategoryList)).Methods("GET")
	r.Handle("/api/goods/list", r.ch.ThenFunc(services.GetGoodsList)).Methods("GET").Queries("k", "{k}").Queries("s", "{s}").Queries("c", "{c}").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/goods/detail", r.ch.ThenFunc(services.GetGoodsDetail)).Methods("GET").Queries("id", "{id}")
	r.Handle("/api/cart/list", r.ch.ThenFunc(services.GetCartGoodsList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/cart/add", r.ch.ThenFunc(services.AddCartGoods)).Methods("POST")
	r.Handle("/api/cart/edit", r.ch.ThenFunc(services.EditCartGoods)).Methods("POST")
	r.Handle("/api/cart/goods_num", r.ch.ThenFunc(services.GetCartGoodsNum)).Methods("GET")
	r.Handle("/api/coupon/list", r.ch.ThenFunc(services.GetCouponList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/coupon/take", r.ch.ThenFunc(services.TakeCoupon)).Methods("POST")
	r.Handle("/api/user/coupon/list", r.ch.ThenFunc(services.GetUserCouponList)).Methods("GET").Queries("status", "{status}").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/user/coupon", r.ch.ThenFunc(services.DoDeleteCouponLog)).Methods("DELETE").Queries("id", "{id}")
	r.Handle("/api/user/address/list", r.ch.ThenFunc(services.GetAddressList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/user/address/edit", r.ch.ThenFunc(services.EditAddress)).Methods("POST")
	r.Handle("/api/user/address", r.ch.ThenFunc(services.GetAddress)).Methods("GET").Queries("id", "{id}")
	r.Handle("/api/user/address", r.ch.ThenFunc(services.DoDeleteAddress)).Methods("DELETE").Queries("id", "{id}")
	r.Handle("/api/user/default_address", r.ch.ThenFunc(services.GetDefaultAddress)).Methods("GET")
	r.Handle("/api/placeorder", r.ch.ThenFunc(services.PlaceOrder)).Methods("POST")
	r.Handle("/api/order/list", r.ch.ThenFunc(services.GetOrderList)).Methods("GET").Queries("status", "{status}").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/order/detail", r.ch.ThenFunc(services.GetOrderDetail)).Methods("GET").Queries("orderNo", "{orderNo}")
	r.Handle("/api/order/cancel", r.ch.ThenFunc(services.CancelOrder)).Methods("PUT").Queries("id", "{id}")
	r.Handle("/api/order", r.ch.ThenFunc(services.DeleteOrder)).Methods("DELETE").Queries("id", "{id}")
	r.Handle("/api/order/confirm_goods", r.ch.ThenFunc(services.ConfirmTakeGoods)).Methods("PUT").Queries("id", "{id}")
	r.Handle("/api/order/refund_apply", r.ch.ThenFunc(services.RefundApply)).Methods("PUT")
	r.Handle("/api/order/refund_detail", r.ch.ThenFunc(services.RefundDetail)).Methods("GET").Queries("refundNo", "{refundNo}")
	r.Handle("/api/order/refund_undo", r.ch.ThenFunc(services.UndoRefundApply)).Methods("PUT").Queries("refundNo", "{refundNo}")
	r.Handle("/api/order/remind", r.ch.ThenFunc(services.GetOrderRemind)).Methods("GET")
	r.Handle("/api/browse/list", r.ch.ThenFunc(services.UserBrowseHistory)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/api/browse/clear", r.ch.ThenFunc(services.ClearBrowseHistory)).Methods("POST")
	r.Handle("/wxpay/notify", r.ch.ThenFunc(services.WxPayNotify)).Methods("POST")
	r.Handle("/cms/user/login", r.ch.ThenFunc(services.CmsUserLogin)).Methods("POST", "OPTIONS")
	r.Handle("/cms/user/refresh", r.ch.ThenFunc(services.Refresh)).Methods("GET", "OPTIONS")
	r.Handle("/cms/user/info", r.ch.ThenFunc(services.GetUserInfo)).Methods("GET", "OPTIONS")
	r.Handle("/cms/user/change_password", r.ch.ThenFunc(services.DoChangePassword)).Methods("PUT", "OPTIONS")
	r.Handle("/cms/admin/users", r.ch.ThenFunc(services.GetUserList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/admin/user", r.ch.ThenFunc(services.DoEditUser)).Methods("POST", "OPTIONS")
	r.Handle("/cms/admin/user", r.ch.ThenFunc(services.GetUser)).Methods("GET", "OPTIONS").Queries("id", "{id}")
	r.Handle("/cms/admin/user", r.ch.ThenFunc(services.DoDeleteCMSUser)).Methods("DELETE", "OPTIONS").Queries("id", "{id}")
	r.Handle("/cms/admin/reset_password", r.ch.ThenFunc(services.DoResetCMSUserPassword)).Methods("POST", "OPTIONS")
	r.Handle("/cms/admin/groups", r.ch.ThenFunc(services.GetUserGroupList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/admin/group", r.ch.ThenFunc(services.DoEditUserGroup)).Methods("POST", "OPTIONS")
	r.Handle("/cms/admin/group", r.ch.ThenFunc(services.GetUserGroup)).Methods("GET", "OPTIONS").Queries("id", "{id}")
	r.Handle("/cms/admin/group", r.ch.ThenFunc(services.DoDeleteUserGroup)).Methods("DELETE", "OPTIONS").Queries("id", "{id}")
	r.Handle("/cms/admin/authority", r.ch.ThenFunc(services.GetModuleList)).Methods("GET", "OPTIONS")
	r.Handle("/cms/banner/list", r.ch.ThenFunc(services.GetBannerList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/banner/{id:[0-9]+}", r.ch.ThenFunc(services.GetBanner)).Methods("GET", "OPTIONS")
	r.Handle("/cms/banner/edit", r.ch.ThenFunc(services.DoEditBanner)).Methods("POST", "OPTIONS")
	r.Handle("/cms/banner/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteBanner)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/category/list", r.ch.ThenFunc(services.GetCategoryList)).Methods("GET", "OPTIONS").Queries("pid", "{pid}").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/category/{id:[0-9]+}", r.ch.ThenFunc(services.GetCategoryById)).Methods("GET", "OPTIONS")
	r.Handle("/cms/category/edit", r.ch.ThenFunc(services.DoEditCategory)).Methods("POST", "OPTIONS")
	r.Handle("/cms/category/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteCategory)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/category/all", r.ch.ThenFunc(services.GetChooseCategory)).Methods("GET", "OPTIONS")
	r.Handle("/cms/grid_category/list", r.ch.ThenFunc(services.GetGridCategoryList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/grid_category/{id:[0-9]+}", r.ch.ThenFunc(services.GetGridCategory)).Methods("GET", "OPTIONS")
	r.Handle("/cms/grid_category/edit", r.ch.ThenFunc(services.DoEditGridCategory)).Methods("POST", "OPTIONS")
	r.Handle("/cms/grid_category/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteGridCategory)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/spec/list", r.ch.ThenFunc(services.GetSpecificationList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/spec/{id:[0-9]+}", r.ch.ThenFunc(services.GetSpecification)).Methods("GET", "OPTIONS")
	r.Handle("/cms/spec/edit", r.ch.ThenFunc(services.DoEditSpecification)).Methods("POST", "OPTIONS")
	r.Handle("/cms/spec/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteSpecification)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/spec/attr/list", r.ch.ThenFunc(services.GetSpecificationAttrList)).Methods("GET", "OPTIONS").Queries("specId", "{specId}")
	r.Handle("/cms/spec/attr/{id:[0-9]+}", r.ch.ThenFunc(services.GetSpecificationAttr)).Methods("GET", "OPTIONS")
	r.Handle("/cms/spec/attr/edit", r.ch.ThenFunc(services.DoEditSpecificationAttr)).Methods("POST", "OPTIONS")
	r.Handle("/cms/spec/attr/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteSpecificationAttr)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/goods/list", r.ch.ThenFunc(services.GetGoodsList)).Methods("GET", "OPTIONS").Queries("k", "{k}").Queries("c", "{c}").Queries("o", "{o}").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/goods/{id:[0-9]+}", r.ch.ThenFunc(services.GetGoods)).Methods("GET", "OPTIONS")
	r.Handle("/cms/goods/edit", r.ch.ThenFunc(services.DoEditGoods)).Methods("POST", "OPTIONS")
	r.Handle("/cms/goods/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteGoods)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/goods/spec", r.ch.ThenFunc(services.GetGoodsSpecList)).Methods("GET", "OPTIONS").Queries("id", "{id}")
	r.Handle("/cms/goods/all", r.ch.ThenFunc(services.GetChooseCategoryGoods)).Methods("GET", "OPTIONS")
	r.Handle("/cms/sku/list", r.ch.ThenFunc(services.GetSKUList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}").Queries("goodsId", "{goodsId}").Queries("k", "{k}").Queries("o", "{o}")
	r.Handle("/cms/sku/{id:[0-9]+}", r.ch.ThenFunc(services.GetSKU)).Methods("GET", "OPTIONS")
	r.Handle("/cms/sku/edit", r.ch.ThenFunc(services.DoEditSKU)).Methods("POST", "OPTIONS")
	r.Handle("/cms/sku/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteSKU)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/coupon/list", r.ch.ThenFunc(services.GetCouponList)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/coupon/{id:[0-9]+}", r.ch.ThenFunc(services.GetCoupon)).Methods("GET", "OPTIONS")
	r.Handle("/cms/coupon/edit", r.ch.ThenFunc(services.DoEditCoupon)).Methods("POST", "OPTIONS")
	r.Handle("/cms/coupon/{id:[0-9]+}", r.ch.ThenFunc(services.DoDeleteCoupon)).Methods("DELETE", "OPTIONS")
	r.Handle("/cms/oss/policy-token", r.ch.ThenFunc(services.GetOSSPolicyToken)).Methods("GET", "OPTIONS").Queries("dir", "{dir}")
	r.Handle("/cms/market_metrics", r.ch.ThenFunc(services.GetMarketMetrics)).Methods("GET", "OPTIONS")
	r.Handle("/cms/order/order_statement", r.ch.ThenFunc(services.GetSaleTableData)).Methods("GET", "OPTIONS").Queries("page", "{page}").Queries("size", "{size}")
	r.Handle("/cms/order/list", r.ch.ThenFunc(services.GetOrderList)).Methods("GET", "OPTIONS").Queries("status", "{status}").Queries("stype", "{stype}").Queries("k", "{k}").Queries("st", "{st}").Queries("et", "{et}").Queries("p", "{p}").Queries("s", "{s}")
	r.Handle("/cms/order", r.ch.ThenFunc(services.GetOrderDetail)).Methods("GET", "OPTIONS").Queries("orderNo", "{orderNo}")
	r.Handle("/cms/order/export", r.ch.ThenFunc(services.ExportOrder)).Methods("GET", "OPTIONS").Queries("status", "{status}").Queries("stype", "{stype}").Queries("k", "{k}").Queries("st", "{st}").Queries("et", "{et}")
	r.Handle("/cms/order/modify_status", r.ch.ThenFunc(services.ModifyOrderStatus)).Methods("PUT", "OPTIONS")
	r.Handle("/cms/order/modify_remark", r.ch.ThenFunc(services.ModifyOrderRemark)).Methods("PUT", "OPTIONS")
	r.Handle("/cms/order/modify_goods", r.ch.ThenFunc(services.ModifyOrderGoods)).Methods("PUT", "OPTIONS")
}
