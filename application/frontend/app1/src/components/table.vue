<!--
 * @Author: Wen Jiajun
 * @Date: 2022-05-08 20:54:19
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 18:17:15
 * @FilePath: \application\frontend\app1\src\components\table.vue
 * @Description: 
-->
<script>
import {ElMessage, ElLoading} from 'element-plus'

export default {
    data(){
        return {
            tableData: [],
            searchTable: [],
            dataView: [],
            dataTitle: '数据',
            dataDialogFormVisible: false,
            form: {
                "table_type": '',
                "file": [],
            },
            searchInput: false,
            idExample: '306d586d-8b49-4117-a287-ab928e73796d',
            searchButton: true,
            publicURL: '/table/public',
            
            fileList: [],
            dialogFormVisible: false,
            fullscreenLoading: false,
        }
    },
    methods:{
        formatData(table){
            var res = []
            for( let i = 0; i < table.labels.length; i++) {
                let tmp = {
                    'label': table.labels[i],
                }
                for ( let c = 0; c < table.columns.length; c++) {
                    tmp[table.columns[c]] = table.data[c][i]
                }
                res.append(tmp)
            }
            return res
        },
        
        start(){
            const loading = ElLoading.service({
                lock: true,
                text: '加载中',
                background: 'rgba(0, 0, 0, 0.7)',
            })         
            this.$axios.get('/tables').then((res) => {
                loading.close()
                console.log(res.data)
                let data = res.data
                if (data['status'] == 200) {
                    this.tableData = data['data']
                    this.searchTable = data['data']
                    ElMessage.success("查询成功")
                } else if (data['status'] > 30000) {
                     ElMessage.info(data['message'])
                } else {
                    ElMessage.error(data['message'])
                }
            }).catch((e) => {
                loading.close()
                console.log(e)
                console.log(this)
                ElMessage.error("暂时无法访问服务器")
            })
        },

        upload(){
            this.$refs.upload.submit()
        },
        
        changeFile(file, fileList){
            console.log("===============")
            console.log(file)
            console.log(fileList)
            this.form.file = fileList[0]
        },

        readMyTableByID(){
            let id = this.searchInput

            // Loading
            const loading = ElLoading.service({
                lock: true,
                text: 'Loading',
                background: 'rgba(0, 0, 0, 0.7)',
            }) 

            this.$axios.get(`/tables/:${id}`).then(res => {
                loading.close()
                let data = res.data
                if (data.status == 200) {
                    this.tableData = data['data']
                    ElMessage.success(data['message'])
                } else if (data['status'] > 30000) {
                     ElMessage.info(data['message'])
                } else {
                    ElMessage.error(data['message'])
                }
            }).catch(e => {
                loading.close()
                console.log(e)
                ElMessage.error("暂时无法连接服务器")
            })
        },

        extract(data){
            return data.toString().slice(0, 10) + '...'
        },

        httpUpload(param) {
            let fd = new FormData()
            console.log('param.file.raw:', param.file)
            fd.append("file", param.file)
            fd.append('table_type', this.form.tableType)
           
            this.dialogFormVisible = false
            const loading = ElLoading.service({
                lock: true,
                text: '数据上链中，请耐心等待',
                background: 'rgba(0, 0, 0, 0.7)',
            })

            console.log(fd)
            console.log(fd.get('file'))

            this.$axios.post('/tables', fd, {headers: {'Content-Type': 'multipart/form-data'}}).then(res => {
                loading.close()
                let data = res.data
                if (data['status'] != 200) ElMessage.error(data['message'])
                else ElMessage.success("上传成功")
            }).catch(e => {
                loading.close()
                console.log(e)
                ElMessage.error("暂时无法连接到服务器")
            })
        },

        viewData(row) {
            let table = row
            this.dataDialogFormVisible = true
            console.log("This Table:", table)
            this.dataView = this.formatData(table)
        },



    },

    created(){
        this.start()
    },

    watch: {
        searchInput(){
            if (this.searchInput.length == 36) {
                this.searchButton = false
            } else {
                this.searchButton = true
                if (!this.searchInput.length) this.start()
            }
        }
    }

}
</script>

<template>
    <el-input 
        class="w-50 m-2" 
        v-model="searchInput" 
        placeholder="输入表格ID进行查询" 
        style="width: 400px"   
        clearable/>  &nbsp 
    <el-button type="primary" @click="getTable()" :disabled="searchButton" round>查询</el-button>
    <el-button type="primary" @click="dialogFormVisible = true" round>新增</el-button>


 <!-- {"id":"1","area":"哈尔滨","year":"2021","month":"10","columns":["装机容量（千瓦）","本月发电量（万千瓦时）","本季累计发电量（万千瓦时）","本月止累计发电量（万千瓦时）","本月上网电量（万千瓦时）","本季累计上网电量（万千瓦时）","本月止累计上网电量（万千瓦时）","本月综合厂用电量（万千瓦时）","本季累计综合厂用电量（万千瓦时）","本月止累计综合厂用电量（万千瓦时）","本月自发自用电量（万千瓦时）","本季累计自发自用电量（万千瓦时）","本月止累计自发自用电量（万千瓦时）","本月其他电量","本月止累计其他电量","本月购网电量","本月止累计购网电量","电厂个数"],"data"  -->
    <el-table :data="searchData" style="width: 100%" fit>
        <el-table-column prop="id" label="表格编号" width="180" > 
            <template #default="scope">
            <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.id }}</span>
            </div>
            </template>
        </el-table-column>
        <el-table-column prop="area" label="地区" width="180" />
        <el-table-column prop="year" label="年份"  width="180"/>
        <el-table-column prop="month" label="月份"  width="180" />
        <el-table-column prop="columns" label="属性"  width="180">
            <template #default="scope">
            <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ extract(scope.row.data) }}</span>

            </div>
 
            </template>
 
        </el-table-column>
        <el-table-column prop="table_type" label="表格类型" width="180"></el-table-column>

        <el-table-column label="操作">
        <template #default="scope">
            <el-button type="success" size="small" @click="viewData(scope.row)" round>查看</el-button>
            <el-button type="primary" size="small" @click="public(scope.row)" round :disabled="scope.row.public == '公开'">公开</el-button>
        </template>
        </el-table-column>
    </el-table>
    
    <el-dialog v-model="dialogFormVisible" title="新增表格">
        <el-form :model="form" label-width="120px">
            <el-form-item label="文件上传">
            <el-upload
            class="upload-demo"
            ref="upload"
            multiple
            :on-change="changeFile"
            :file-list="fileList"
            action=""
            :limit="3"
            :auto-upload="false"
            :http-request="httpUpload"
            >
            <el-button type="primary" round><el-icon><upload /></el-icon>选择文件</el-button>
            </el-upload>
        </el-form-item>
        <el-form-item label="数据公开">
            <el-select v-model="form.public" placeholder="请选择是否加密数据">
                <el-option label="加密" value="false" />
                <el-option label="不加密" value="true" />
            </el-select>
        </el-form-item>
        <!-- ================================================= -->
        <el-form-item label="备注">
        <el-input v-model="form.description" type="textarea" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" 
                @click="upload()" 
                :disabled="fileList.length === 0" 
                >上传</el-button>
            <el-button @click="dialogFormVisible = false">取消</el-button>
        </el-form-item>
        </el-form>
    </el-dialog>

<!-- columns labels data [][]float64 -->
    <el-dialog v-model="dataDialogFormVisible" title="具体数据">
        <el-table :data="dataView" style="width: 100%" fit>


            <template #default="scope">
                <el-table-column prop="label" label="观测" width="180" /> 
                <el-table-column v-for="col in scope.row" :prop="col" :label="col" width="180" />
            </template>
        </el-table>
    </el-dialog>

</template>


