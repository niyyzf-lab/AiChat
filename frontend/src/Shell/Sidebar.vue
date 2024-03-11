<script setup>
import "../minx.scss"
import Logo from "../components/Logo.vue";
</script>

<template>
  <div class="Box">
    <div class="Box2">
      <div class="Menu1">
        <Logo/>
        <div>
          <span>
            <img alt="Add" src="../assets/images/add.svg"/>
            <span>New chat</span>
          </span>
          <span>
            <img alt="Add" src="../assets/images/search-normal.svg"/>
          </span>
        </div>
      </div>
      <div class="HistoricalList">
        <div class="Header">
          <span>Your conversations</span>
          <span>Clear All</span>
        </div>
        <div class="HistoryContent">

          <router-link to="/">Home</router-link>
          <router-link to="/chat">Chat</router-link>

        </div>
      </div>
      <div class="Menu2" @click="$router.push('/setting')">
        <span>
          <img src="../assets/images/setting-2.svg" alt="setting"/>
        </span>
        <p>Settings</p>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.Box {
  width: 400px;
  height: 100%;
  padding:  20px;
  box-sizing: border-box;

  .Box2 {
    background: white;
    border-radius: 20px;
    box-shadow: 0 8px 24px -4px rgba(0, 0, 0, 0.08);
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;;
  }
}

.Menu1 {
  padding: 30px;

  & > div {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    margin-top: 50px;

    & > span:nth-child(1) {
      display: flex;
      flex: 1;
      height: 50px;
      border-radius: 69px;
      background: #5661F6;
      align-items: center;
      justify-content: center;
      gap: 10px;

      span {
        color: #FFF;
        font-size: 16px;
        font-style: normal;
        font-weight: 500;
        line-height: 22px; /* 137.5% */
      }
    }

    & > span:nth-child(2) {
      height: 50px;
      width: 50px;
      border-radius: 69px;
      background: #000;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}

.HistoricalList {
  display: flex;
  flex-direction: column;
  flex: 1;
  .Header {
    display: flex;
    justify-content: space-between;
    padding: 30px;
    //y轴边框
    border: {
      bottom: 1px solid #E5E5E5;
      top: 1px solid #E5E5E5;
    };

    span {
      font-size: 14px;
      font-style: normal;
      line-height: 20px; /* 142.857% */
    }

    span:nth-child(1) {
      color: #6A6969;
      font-weight: 500;
    }

    span:nth-child(2) {
      color: #5661F6;
      font-weight: 600;

      &:hover {
        cursor: pointer;
        text-decoration: underline;
      }
    }
  }
  .HistoryContent{
    flex:1;
  }
}
.Menu2{
  cursor: pointer;
  height: 50px;
  border-radius: 25px;
  border: 1px solid #EFEFEF;
  margin-bottom: 50px;
  margin-inline: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding-left: 8px;
  &>span{
    height: 34px;
    width: 34px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #c1c1c1;
    border-radius: 100px;
  }
  &>p{
    color: #000;
    font-size: 16px;
    font-style: normal;
    font-weight: 500;
    line-height: 20px; /* 125% */
    flex: 1;
  }
}
</style>