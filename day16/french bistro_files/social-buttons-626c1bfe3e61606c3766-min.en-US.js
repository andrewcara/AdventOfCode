(self.webpackChunkextract_css=self.webpackChunkextract_css||[]).push([[62318],{483867:(c,f)=>{"use strict";Object.defineProperty(f,"__esModule",{value:!0}),f.default=void 0;var n;(function(i){i.TEXT=1,i.IMAGE=2,i.QUOTE=4,i.LINK=5,i.CHAT=6,i.AUDIO=7,i.VIDEO=8,i.VIDEO_DEPRECATED=9,i.REVIEW=10,i.STORE_ITEM=11,i.EVENT=12,i.THREAD=13,i.GALLERY=14,i.BINARY=15,i.CSSASSET=16,i.TWEAKASSET=17,i.DIGITALGOOD=18,i.ATTACHMENT=19,i.EXPORT_WORDPRESS=20,i.EXPORT_INTERNAL=21,i.TEXT_PROSE_MIRROR=22,i.EXPORT_ROLES=23,i.TWEET=50,i.RSS=51,i.CHECKIN=52,i.DELICIOUS=53,i.KBARTICLE=54,i.PROJECT_ITEM=55,i.COLLECTION_TAXONOMY=56,i.SECTION_TAXONOMY=57,i.ITEM_TAXONOMY=58,i.PORTFOLIO_ITEM=59,i.EXPORT_TRANSLATABLE_STRINGS=60,i.SQSP_VIDEO=61,i.LESSON=62,i.COURSE_ITEM=63})(n||(n={}));var l=n;f.default=l,c.exports=f.default},50719:(c,f)=>{"use strict";Object.defineProperty(f,"__esModule",{value:!0}),f.default=void 0;var n;(function(i){i.TWITTER=1,i.FACEBOOK=2,i.GOOGLE=3,i.LINKEDIN=4,i.STUMBLE=5,i.REDDIT=6,i.PINTEREST=7,i.TUMBLR=8})(n||(n={}));var l=n;f.default=l,c.exports=f.default},974458:(c,f,n)=>{"use strict";var l=n(883644),i=n(436032),h=n(685241),p=n.n(h)},436032:(c,f,n)=>{"use strict";n.r(f);var l=n(574275),i=n.n(l),h=n(215575),p=n.n(h),b=n(50719),v=n.n(b),d=n(311064),T=n.n(d);YUI.add("squarespace-social-button",function(e){e.namespace("Squarespace");var u=v(),I=400,_=function(s){return s},a=e.Squarespace.SocialButton=e.Base.create("socialButton",e.Widget,[],{initializer:function(s){this._servicesRendered={},s.services.forEach(function(t){this._servicesRendered[t]=!1},this),this._open=!1,this._anims={},this._serviceRenderers={},this._serviceRenderers[u.REDDIT]=this._renderReddit,this._serviceRenderers[u.FACEBOOK]=this._renderFacebook,this._serviceRenderers[u.TWITTER]=this._renderTwitter,this._serviceRenderers[u.LINKEDIN]=this._renderLinkedIn,this._serviceRenderers[u.STUMBLE]=this._renderStumble,this._serviceRenderers[u.PINTEREST]=this._renderPinterest,this._serviceRenderers[u.TUMBLR]=this._renderTumblr,this.publish("serviceRendered",{defaultFn:this._defaultServiceRendered,context:this}),this.publish("buttonClicked",{defaultFn:this._defaultButtonClicked,preventable:!0,context:this}),this.publish("close",{defaultFn:this.close,preventable:!0,context:this}),this.publish("servicesRendered"),this._serviceContainer=e.Node.create('<div class="ss-social-button-container"></div>')},_defaultServiceRendered:function(s){var t=s.details[0];this._servicesRendered[t]=!0,this._allServicesRendered()&&(this.set("loaded",!0),this.fire("servicesRendered",this))},open:function(){this._open=!0,this._openList()},close:function(){this._open=!1,this._closeList()},_onClick:function(s){this.fire("buttonClicked",s)},_defaultButtonClicked:function(s){this.get("loaded")?this.isOpen()?this.close():this.open():this.get("loading")||(this.once("servicesRendered",function(){this.set("loading",!1)},this),this.set("loading",!0),this._renderServices(),this.open())},isOpen:function(){return this._open},destructor:function(){this._stopAnimations()},_stopAnimations:function(){e.Object.values(this._anims).forEach(function(s){s.stop(),s=null})},_closeList:function(){var s=this.get("contentBox");if(s._node&&s.inDoc()){var t=s.one(".ss-social-list-wrapper"),o;t&&t._node&&t.inDoc()&&(o=new e.Anim({node:t,duration:.3,easing:e.Easing.easeOutStrong,to:{height:0,opacity:0}}),this._anims.close=o,o.on("end",function(){t.setStyle("overflow",null),this.fire("listClose")},this),o.run())}},_openList:function(){var s=this.get("contentBox");if(s._node&&s.inDoc()){var t=s.one(".ss-social-button-wrapper"),o=s.one(".ss-social-list-wrapper"),r=s.one(".ss-social-button-list"),E=t.get("offsetWidth"),O=t.get("offsetHeight"),S=o.get("offsetWidth"),g=Math.abs(E-S),L,R;o.setStyles({left:(E<=S?-1:1)*g/2,top:O});var C=e.DOM.viewportRegion(),m=r.get("region"),M=m.height+m.top,P=M-(C.height+C.top),B=C.left-m.left,U=B>0,D=20,N=20;P>0&&o.setStyle("top",parseInt(o.getComputedStyle("top"),10)-P-D),U&&o.setStyle("left",parseInt(o.getComputedStyle("left"),10)+B+N),o&&o._node&&o.inDoc()&&(R=new e.Anim({node:o,duration:.3,easing:e.Easing.easeOutStrong,to:{height:r.get("offsetHeight"),opacity:1}}),this._anims.open=R,R.on("end",function(){o.setStyle("overflow","visible");var A=e.config.win.document;e.UA.touchEnabled&&e.one(A.documentElement).setStyle("cursor","pointer"),L=e.one(A).on("click",function(w){w.target.ancestor(".ss-social-list-wrapper",!0)||(this.fire("close"),L.detach(),L=null,e.one(A.documentElement).setStyle("cursor",""))},this),this.fire("listOpen")},this),R.run())}},_allServicesRendered:function(){return e.Object.values(this._servicesRendered).every(_)},bindUI:function(){var s=this.get("contentBox"),t=s.one(".ss-social-button-wrapper");this.after("loadingChange",this._onLoadingChange,this),t.on("click",this._onClick,this)},_onLoadingChange:function(){this.get("boundingBox").ancestor(".squarespace-social-buttons").toggleClass("loading",this.get("loading"))},renderUI:function(){var s=this.get("contentBox");s.append('<div class="ss-social-button-wrapper"><div class="ss-social-button"><span class="ss-social-button-icon"></span>'+(0,l.t)("Share",null,{project:"scripts-v6-root"})+"</div></div>"),s.append('<div class="ss-social-list-wrapper"><div class="ss-social-button-list"></div></div>')},_renderServices:function(){var s=this.get("contentBox").one(".ss-social-button-list");this.get("services").forEach(function(t){t in this._serviceRenderers&&this._serviceRenderers[t].call(this,s)},this)},_defaultTimeoutCb:function(s,t){var o=T()(v()),r=(0,l.t)("(Social Button) Loading render script for service: {name} too longer than {sub1} seconds. Skipping.",{sub1:a.SCRIPT_TIMEOUT/1e3},{project:"scripts-v6-root"});return function(){this.fire("serviceRendered",s),t&&t.hide()}},_defaultFailureCb:function(s,t){var o=T()(v()),r=(0,l.t)("(Social Button) Service {name} render script failed to load.",null,{project:"scripts-v6-root"});return function(){this.fire("serviceRendered",s),t&&t.hide()}},_renderReddit:function(s){var t=this._serviceContainer.cloneNode(!0),o=this.get("url");t.addClass("reddit"),t.append(e.Node.create(a.REDDIT_LINK)),t.on("click",function(r){window.open("http://reddit.com/submit?url="+encodeURIComponent(o),(0,l.t)("Submit to Reddit",null,{project:"scripts-v6-root"})),r.stopImmediatePropagation()}),s.append(t),this.fire("serviceRendered",u.REDDIT)},_renderTumblr:function(s){var t=this._serviceContainer.cloneNode(!0),o={url:this.get("url"),name:this.get("title")},r=e.QueryString.stringify(o);t.addClass("tumblr"),t.append(e.substitute(a.TUMBLR_TAG_TEMPLATE,{query:r})),s.append(t),e.Get.script(a.TUMBLR_URL,{onSuccess:function(){e.later(I,this,function(){this.fire("serviceRendered",u.TUMBLR)})},onFailure:this._defaultFailureCb(u.TUMBLR,t),onTimeout:this._defaultTimeoutCb(u.TUMBLR,t),timeout:a.SCRIPT_TIMEOUT,context:this,win:e.config.win})},_renderFacebook:function(s){var t=this._serviceContainer.cloneNode(!0),o=this.get("url");t.addClass("facebook"),t.append(e.Node.create(e.substitute(a.FACEBOOK_TAG_TEMPLATE,{url:o}))),s.append(t);var r=e.config.win,E=r.Static.SQUARESPACE_CONTEXT.facebookAppId,O=r.Static.SQUARESPACE_CONTEXT.facebookApiVersion,S=e.substitute(a.FACEBOOK_URL,{locale:h.formatLocaleForFacebook(h.getResolvedWebsiteLocale())});e.Get.script(S,{onSuccess:function(){r.FB&&e.later(I,this,function(){r.FB.init({appId:E,xfbml:!1,version:O}),r.FB.XFBML&&r.FB.XFBML.parse&&r.FB.XFBML.parse(),this.fire("serviceRendered",u.FACEBOOK)})},onFailure:this._defaultFailureCb(u.FACEBOOK,t),onTimeout:this._defaultTimeoutCb(u.FACEBOOK,t),timeout:a.SCRIPT_TIMEOUT,context:this})},_renderTwitter:function(s){var t=this._serviceContainer.cloneNode(!0),o=this.get("title"),r=this.get("url");t.append(e.Node.create('<a href="https://twitter.com/share" data-text="'+e.Escape.html(o||"")+'" data-url="'+e.Escape.html(r)+'"class="twitter-share-button">'+(0,l.t)("tweet",null,{project:"scripts-v6-root"})+"</a>")),t.addClass("twitter"),s.append(t),e.Get.script(a.TWITTER_URL,{onSuccess:function(){e.later(I,this,function(){this.fire("serviceRendered",u.TWITTER)})},onFailure:this._defaultFailureCb(u.TWITTER,t),onTimeout:this._defaultTimeoutCb(u.TWITTER,t),timeout:a.SCRIPT_TIMEOUT,context:this})},_renderLinkedIn:function(s){var t=this._serviceContainer.cloneNode(!0),o=this.get("url");t.addClass("linkedin"),t.append(e.Node.create(e.substitute(a.LINKEDIN_URL_TEMPLATE,{url:o}))),s.append(t),window.IN=void 0,e.Get.script(a.LINKEDIN_URL,{onSuccess:function(){e.later(I,this,function(){this.fire("serviceRendered",u.LINKEDIN)})},onFailure:this._defaultFailureCb(u.LINKEDIN,t),onTimeout:this._defaultTimeoutCb(u.LINKEDIN,t),timeout:a.SCRIPT_TIMEOUT,context:this})},_renderStumble:function(s){var t=this._serviceContainer.cloneNode(!0),o=this.get("url"),r=this.get("id");t.addClass("stumble"),t.append(e.Node.create(e.substitute(a.STUMBLE_TAG_TEMPLATE,{url:o,id:r}))),s.append(t),e.Get.script(a.STUMBLE_URL,{onSuccess:function(){e.later(I,this,function(){window.STMBLPN&&(window.STMBLPN.wasProcessLoaded&&(window.STMBLPN.wasProcessLoaded=!1),window.STMBLPN.processWidgets()),this.fire("serviceRendered",v().STUMBLE)})},onFailure:this._defaultFailureCb(u.STUMBLE,t),onTimeout:this._defaultTimeoutCb(u.STUMBLE,t),timeout:a.SCRIPT_TIMEOUT,context:this})},_renderPinterest:function(s){var t=this._serviceContainer.cloneNode(!0),o=this.get("assetUrl"),r=this.get("url"),E=Static.SQUARESPACE_CONTEXT.website.authenticUrl+s.ancestor(".squarespace-social-buttons").getAttribute("data-full-url");this.get("systemDataId")?(t.addClass("pinterest"),t.append(e.Node.create(e.substitute(a.PINTEREST_TAG_TEMPLATE,{url:encodeURIComponent(o||r),pageUrl:encodeURIComponent(E)}))),s.append(t),e.Get.script(a.PINTEREST_URL,{onSuccess:function(){e.later(I,this,function(){this.fire("serviceRendered",v().PINTEREST)},this)},onFailure:this._defaultFailureCb(u.PINTEREST,t),onTimeout:this._defaultTimeoutCb(u.PINTEREST,t),timeout:a.SCRIPT_TIMEOUT,context:this})):this.fire("serviceRendered",u.PINTEREST)}},{TWITTER_URL:"//platform.twitter.com/widgets.js",TUMBLR_URL:"//platform.tumblr.com/v1/share.js",FACEBOOK_URL:"//connect.facebook.net/{locale}/sdk.js",LINKEDIN_URL:"//platform.linkedin.com/in.js",STUMBLE_URL:"http://platform.stumbleupon.com/1/widgets.js",PINTEREST_URL:"//assets.pinterest.com/js/pinit.js",LINKEDIN_URL_TEMPLATE:'<script type="IN/Share" data-url="{url}" data-counter="right"><\/script>',FACEBOOK_TAG_TEMPLATE:'<div id="fb-root"></div><fb:like href="{url}" send="false" layout="button_count" show_faces="true"></fb:like>',PINTEREST_TAG_TEMPLATE:'<a href="//pinterest.com/pin/create/button?url={pageUrl}&media={url}" class="pin-it-button" count-layout="horizontal"><img border="0" src="//assets.pinterest.com/images/PinExt.png" title="'+(0,l.t)("Pin It",null,{project:"scripts-v6-root"})+'" /></a>',TUMBLR_TAG_TEMPLATE:'<a href="https://tumblr.com/share/link?{query}" title="'+(0,l.t)("Share on Tumblr",null,{project:"scripts-v6-root"})+`" style="display:inline-block; text-indent:-9999px; overflow:hidden; width:81px; height:20px; background:url('https://platform.tumblr.com/v1/share_1T.png') top left no-repeat transparent;">`+(0,l.t)("Share on Tumblr",null,{project:"scripts-v6-root"})+"</a>",SCRIPT_TIMEOUT:5e3,STUMBLE_TAG_TEMPLATE:'<su:badge layout="1" location="{url}"></su:badge>',REDDIT_LINK:'<a href="#"><img src="https://old.reddit.com/static/spreddit7.gif" alt="'+(0,l.t)("submit to reddit",null,{project:"scripts-v6-root"})+'" border="0" /></a>',ATTRS:{url:{value:window.location.href},title:{value:document.title||window.location.href},services:{},recordType:{},assetUrl:{value:""},systemDataId:{value:""},loaded:{value:!1},loading:{value:!1}}})},"1.0",{requires:["anim","base","escape","node","querystring-stringify","squarespace-util","substitute","widget"]})},883644:(c,f,n)=>{"use strict";n.r(f);var l=n(50719),i=n.n(l),h=n(483867),p=n.n(h),b=n(87421),v=n.n(b);YUI.add("squarespace-social-buttons",function(d){d.namespace("Squarespace");var T=d.config.win.Static,e=d.Squarespace.SocialButton,u=d.Squarespace.SocialButtons=d.Base.create("socialButtons",d.Base,[],{initializer:function(_){var a=this.get("services");this._buttonConfigs={},this._buttons=[],this._eventHandles=[],a.length>0?(this._scanForButtons(),this._initializeButtons(),this._bindEvents(),this._renderButtons()):this._markButtonsAsEmpty()},_markButtonsAsEmpty:function(){d.all(u.SOCIAL_BUTTON_CONTAINER).addClass("empty")},_scanForButtons:function(){var _=d.all(u.SOCIAL_BUTTON_CONTAINER),a=this.get("services");_.each(function(s){var t=s.getAttribute(u.TITLE),o=T.SQUARESPACE_CONTEXT.website.baseUrl+s.getAttribute(u.FULL_URL),r=parseInt(s.getAttribute(u.RECORD_TYPE),10),E=s.getAttribute(u.ASSET_URL),O=s.getAttribute(u.SYSTEM_DATA_ID),S=d.guid(u.ID_PREFIX);this._buttonConfigs[S]={id:S,title:t,url:o,recordType:r,assetUrl:E,systemDataId:O,services:a,node:s}},this)},_initializeButtons:function(){this._buttons=d.Array.map(d.Object.values(this._buttonConfigs),function(_){return new e(_)})},_bindEvents:function(){this._eventHandles.push(this.after("cleanup",this._defaultDestroy,this))},_renderButtons:function(){this._buttons.filter(this._excludeOnlyPinterest,this).forEach(function(_){var a=_.get("id");_.render(this._buttonConfigs[a].node)},this)},destructor:function(){this.fire("cleanup")},_unbindEvents:function(){this._eventHandles.forEach(function(_){_.detach(),_=null})},_defaultDestroy:function(){this._destroyButtons(),this._unbindEvents(),this.fire("destroyed")},_destroyButtons:function(){this._buttons.forEach(function(_){_.destroy()},this)},_excludeOnlyPinterest:function(_){var a=_.get("recordType");return!(a!==p().IMAGE&&this._onlyServiceIsPinterest(_))},_onlyServiceIsPinterest:function(_){var a=_.get("services");return a.length===1&&a[0]===i().PINTEREST}},{FULL_URL:"data-full-url",ASSET_URL:"data-asset-url",SYSTEM_DATA_ID:"data-system-data-id",RECORD_TYPE:"data-record-type",ID_PREFIX:"social-",TITLE:"data-title",SOCIAL_BUTTON_CONTAINER:".squarespace-social-buttons",ATTRS:{services:{valueFn:function(){return d.Array.map(d.Object.keys(T.SQUARESPACE_CONTEXT.website.shareButtonOptions||{}),function(_){return parseInt(_,10)})}}}}),I=[];d.config.win.Squarespace.onInitialize(d,function(){d.all(".squarespace-social-buttons").isEmpty()||I.push(new d.Squarespace.SocialButtons)}),d.config.win.Squarespace.onDestroy(d,function(){I.forEach(function(_){_.destroy()}),I.length=0})},"1.0",{requires:["array-extras","base","node","squarespace-social-button"]})},431345:(c,f,n)=>{var l=n(955232);function i(h,p,b,v){return l(h,function(d,T,e){p(v,b(d),T,e)}),v}c.exports=i},614661:(c,f,n)=>{var l=n(431345);function i(h,p){return function(b,v){return l(b,h,p(v),{})}}c.exports=i},764483:c=>{function f(n){return function(){return n}}c.exports=f},311064:(c,f,n)=>{var l=n(764483),i=n(614661),h=n(839039),p=Object.prototype,b=p.toString,v=i(function(d,T,e){T!=null&&typeof T.toString!="function"&&(T=b.call(T)),d[T]=e},l(h));c.exports=v},87421:c=>{c.exports={}},685241:()=>{YUI.add("substitute",function(c,f){var n=c.Lang,l="dump",i=" ",h="{",p="}",b=/(~-(\d+)-~)/g,v=/\{LBRACE\}/g,d=/\{RBRACE\}/g,T=function(e,u,I,_){for(var a,s,t,o,r,E,O=[],S,g,L=e.length;a=e.lastIndexOf(h,L),!(a<0||(s=e.indexOf(p,a),a+1>=s));)S=e.substring(a+1,s),o=S,E=null,t=o.indexOf(i),t>-1&&(E=o.substring(t+1),o=o.substring(0,t)),r=u[o],I&&(r=I(o,r,E)),n.isObject(r)?c.dump?n.isArray(r)?r=c.dump(r,parseInt(E,10)):(E=E||"",g=E.indexOf(l),g>-1&&(E=E.substring(4)),r.toString===Object.prototype.toString||g>-1?r=c.dump(r,parseInt(E,10)):r=r.toString()):r=r.toString():n.isUndefined(r)&&(r="~-"+O.length+"-~",O.push(S)),e=e.substring(0,a)+r+e.substring(s+1),_||(L=a-1);return e.replace(b,function(R,C,m){return h+O[parseInt(m,10)]+p}).replace(v,h).replace(d,p)};c.substitute=T,n.substitute=T},"3.17.2",{requires:["yui-base"],optional:["dump"]})},392338:c=>{"use strict";c.exports=void 0}},c=>{var f=l=>c(c.s=l);c.O(0,[80276,46001],()=>f(974458));var n=c.O()}]);

//# sourceMappingURL=https://sourcemaps.squarespace.net/universal/scripts-compressed/sourcemaps/8f6d094b47ed772b7c92005b6e123d5e/social-buttons-626c1bfe3e61606c3766-min.en-US.js.map