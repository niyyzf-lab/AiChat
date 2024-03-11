<script setup>
import {useRouter} from "vue-router";
import {PostOllma} from "../../wailsjs/go/main/App.js";
const theInputBoxIsDisplayed = ref(true);
const route = useRouter()
//当页面更新时获取当前路由
route.afterEach((to, from) => {
  const currentPath = route.currentRoute.value.path
  //如果是setting页面则设为none
  theInputBoxIsDisplayed.value = currentPath !== '/setting';
});
const InputValue= ref("");
let InputData=[];
const SendClick=(event)=>{
  event.preventDefault()
  const str=InputValue.value
  if(str.length>0){
    const inStr={
      "role":"user",
      "content":str
    }
    InputData=[inStr,...InputData]
    PostOllma("gemma",JSON.stringify(InputData)).then(
        (data)=>{
          console.log(data)
          const message={
            "role":"assistant",
            "content":data
          }
          InputData=[message,...InputData]
        }
    )
    InputValue.value=""
  }else{
    alert("输入不能为空")
  }
}

</script>

<template>
<div class="Box1">
  <div id="ViewsBox">
    <keep-alive>
      <router-view>
      </router-view>
    </keep-alive>
  </div>
  <form class="InputBar" @submit="SendClick" v-show="theInputBoxIsDisplayed">
    <img src="./../assets/images/naozi.png" alt="naozi">
    <input placeholder="What's in your mind?..." type="text" v-model="InputValue">
    <button type="submit">
      <img src="../assets/images/send-2.svg" alt="">
    </button>
  </form>
</div>
</template>

<style scoped lang="scss">
.Box1{
  flex: 1;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
}
#ViewsBox{
  overflow-y: auto;
  display: flex;
  width: 100%;
  height: 100%;
}
.InputBar{
  position: absolute;
  bottom: 60px;
  height: 60px;
  display: flex;
  width: 747px;
  padding: 6px 6px 6px 20px;
  justify-content: space-between;
  align-items: center;
  border-radius: 30px;
  background: #FFF;
  box-shadow: 0 8px 24px -4px rgba(0, 0, 0, 0.08);
  &>button{
    display: flex;
    width: 48px;
    height: 48px;
    padding: 10px;
    justify-content: center;
    align-items: center;
    gap: 10px;
    flex-shrink: 0;
    border-radius: 60px;
    background: #5661F6;
    box-shadow: 0 4px 8px 0 rgba(86, 97, 246, 0.25);
  }
  &>input{
    flex: 1;
    margin-inline: 10px;
    &:focus{
      outline: none;
    }
  }
}
</style>