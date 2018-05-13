var apiUrl = "/pullrequests";

var uri = window.location.protocol + '//' + window.location.host + apiUrl + window.location.search;

Vue.component("repository", {
    props: ["repo"],
    template: `
    <div class="repo">
        <a v-bind:href="repo.url">
            <span class="repo-owner">{{ repo.Owner.Login }}/</span>
            <span class="repo-name">{{ repo.Name }}</span>
        </a>
        <div class="repo-prs">

            <pullrequest v-for="pr in repo.PullRequests.Nodes" :key="pr.url" v-bind:pr="pr"></pullrequest>
        </div>
    </div>
    `
});

Vue.component("pullrequest", {
    props: ["pr"],
    template: `
    <div class="pr">
        <div class="pr-details">
            <svg class="pr-icon" viewBox="0 0 12 16" version="1.1" width="24" height="24" aria-hidden="true">
                <path class="pr-icon" fill-rule="evenodd" d="M11 11.28V5c-.03-.78-.34-1.47-.94-2.06C9.46 2.35 8.78 2.03 8 2H7V0L4 3l3 3V4h1c.27.02.48.11.69.31.21.2.3.42.31.69v6.28A1.993 1.993 0 0 0 10 15a1.993 1.993 0 0 0 1-3.72zm-1 2.92c-.66 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2zM4 3c0-1.11-.89-2-2-2a1.993 1.993 0 0 0-1 3.72v6.56A1.993 1.993 0 0 0 2 15a1.993 1.993 0 0 0 1-3.72V4.72c.59-.34 1-.98 1-1.72zm-.8 10c0 .66-.55 1.2-1.2 1.2-.65 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2zM2 4.2C1.34 4.2.8 3.65.8 3c0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2z"></path>
            </svg>
            <img class="pr-avatar" v-bind:src="pr.Author.AvatarURL">
            <span class="pr-id">#{{ pr.Number }}</span>
            <div class="pr-title">{{ pr.Title }}</div>
            <div class="pr-difficulty">{{ pr.Difficulty }}</div>
            <div class="pr-lines">
                <div class="pr-additions">+{{ pr.Additions }}</div>
                <div class="pr-deletions">-{{ pr.Deletions }}</div>
                <div class="pr-files">{{ pr.ChangedFiles }} file(s)</div>
            </div>
            <div class="pr-date">
                {{ pr.Since }}
            </div>
        </div>
        <div class="pr-comments">
            <div class="pr-comment" v-for="c in pr.Comments.Nodes" :key="c.URL">
                <img class="comment-avatar" v-bind:src="c.Author.AvatarURL">
                <span class="comment-text">{{ c.BodyText }}</span>
            </div>
        </div>
    </div>
    `
});

var weight = {
    "addition": 1,
    "deletion": 0.7,
    "file": 2
};

var difficulty = weight => {
    if (weight < 15) {
        return "LIGHTWEIGHT";
    } else if (weight < 70) {
        return "SUPER EASY";
    } else if (weight < 150) {
        return "EASY";
    } else if (weight < 350) {
        return "MEDIUM;"
    } else if (weight < 600) {
        return "HARD";
    } else {
        return "DIFFICULT";
    }
}

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
                .then(result => {
                    result.forEach(repo => {
                        repo.PullRequests.Nodes.forEach(pr => {
                            var now = moment();
                            var prCreated = moment(pr.CreatedAt);
                            pr.Since = now.from(prCreated, true);

                            var w = pr.Additions * weight.addition +
                                pr.Deletions * weight.deletion +
                                pr.ChangedFiles * weight.file;

                            pr.Difficulty = difficulty(w);
                        });
                    });
                    this.repos = result
                });
        }
    },
    mounted() {
        return this.getData()
    }
});
