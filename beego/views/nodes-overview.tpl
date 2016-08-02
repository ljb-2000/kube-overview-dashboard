<a href="/">retour</a>
<!-- /.row -->
<div class="row">
    <div class="col-lg-12">
        <div class="panel panel-default">
            <div class="panel-heading">
                {{ .title }}
            </div>
            <!-- /.panel-heading -->
            <div class="panel-body">
                <div class="dataTable_wrapper">
                    <table class="table table-striped table-bordered table-hover" id="dataTables-example">
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>Label</th>
                            </tr>
                        </thead>
                        <tbody>
                          {{ range $key, $value := .nodes }}
                            <tr class="odd gradeX">
                                <td>{{ $key }}</td>
                                <td>{{ $value.labels }}</td>
                            </tr>
                            {{ end }}
                        </tbody>
                    </table>
                </div>
            </div>
            <!-- /.panel-body -->
        </div>
        <!-- /.panel -->
    </div>
    <!-- /.col-lg-12 -->
</div>
