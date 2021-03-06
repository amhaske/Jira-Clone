import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, FormArray} from '@angular/forms';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute, Router } from '@angular/router';
import { ProjectSettings, project_data } from '../project-list/project-list.component';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-project-settings,',
  templateUrl: './project-settings.component.html',
  styleUrls: ['./project-settings.component.scss']
})
export class ProjectSettingsComponent implements OnInit {
  
  project_name = ""
  project_desc = ""
  project_key = ""
  project_id = ""
  created_at = ""
  owner_username = ""
  memberCnt = ""
  issueCnt = ""
  commentCnt = ""

  constructor(private fb: FormBuilder, private apiService:ApiInterfaceService, private route: ActivatedRoute, private router: Router, private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.get_project_data();  
     
    // set the form values
    // this.setFormValues();

  }

  go_to_sprints(): void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.router.navigateByUrl('home/project/'+projectIdFromRoute+'/sprints')

  }

  get_project_data():void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'))
    this.apiService.getProjectDetails({'project_id':projectIdFromRoute})
    .subscribe((resp:any) => {
      console.log(resp['resp']);
      this.project_name = resp['resp']['project_name'];
      this.project_desc = resp['resp']['project_desc'];
      this.project_key = resp['resp']['project_key'];
      this.project_id = resp['resp']['project_id'];
      this.created_at = resp['resp']['created_at'];
      this.owner_username = resp['resp']['owner_uname'];
      this.get_project_stats(); 
    })
  }

  get_project_stats(){
    const body = {
      "project_id": this.project_id
    }
    this.apiService.getProjectStats(body)
    .subscribe((resp:any) => {
      console.log(resp['resp']);
      this.memberCnt = resp['resp']['member_count'];
      this.issueCnt = resp['resp']['issue_count'];
      this.commentCnt = resp['resp']['comment_count'];
    })

  }

  delete_project():void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));

    let body = {
      "user_id": 1,
      "project_id": projectIdFromRoute,
    }

    this.apiService.deleteProject(body)
    .subscribe((resp:any) =>{
      console.log(body)
      console.log(resp['response'])
      this.createAlert("Project deleted successfully!")
      this.router.navigateByUrl('/home/projects');
    })
  }


  createAlert(message:string): void{
    this._snackBar.open(message, "Done");
  }

}
