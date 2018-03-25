var apiUrl = "/pullrequests";

var uri = window.location.protocol + '//' + window.location.host + apiUrl + window.location.search;

Vue.component("repository", {
    props: ["repo"],
    template: `
    <div class="repo">
        <span class="repo-owner">{{ repo.owner }}/</span>
        <span class="repo-name">{{ repo.name }}</span>
        <div class="repo-prs">
            <pullrequest v-for="pr in repo.pull_requests" :key="pr.url" v-bind:pr="pr"></pullrequest>
        </div>
    </div>
    `
});

Vue.component("pullrequest", {
    props: ["pr"],
    template: `
    <div class="pullrequest">
        <span class="pr-title"> >>> {{ pr.title }}/</span>
        <img class="pr-avatar" v-bind:src="pr.user.avatar_url">
    </div>
    `
});

var app = new Vue({
    el: "#app",
    data() {
        return {
            repos: []
        }
    },
    methods: {
        getData() {
            return fetch(uri, {
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json"
                }
            })
            .then(response => response.json())
            .then(result => this.repos = result);
        }
    },
    mounted() {
        return this.getData()
    }
});